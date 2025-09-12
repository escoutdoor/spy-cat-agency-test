package target

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	def "github.com/escoutdoor/spy-cat-agency-test/internal/repository"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/database"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type repository struct {
	db database.Client
}

var _ def.TargetRepository = (*repository)(nil)

func New(db database.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) UpdateTarget(ctx context.Context, in dto.UpdateTargetParams) error {
	sql := `
		UPDATE targets
		SET
	`

	args := pgx.NamedArgs{}
	var updates []string

	if in.Completed != nil {
		args["completed"] = in.Completed
		updates = append(updates, "completed=@completed")
	}
	if in.Notes != nil {
		args["notes"] = in.Notes
		updates = append(updates, "notes=@notes")
	}

	if len(updates) == 0 {
		return apperrors.NoFieldsToUpdate
	}

	sql += fmt.Sprintf("updated_at=now(), %s WHERE id=@target_id", strings.Join(updates, ", "))
	args["target_id"] = in.ID

	q := database.Query{
		Name: "target_repository.UpdateTarget",
		Sql:  sql,
	}
	if _, err := r.db.DB().ExecContext(ctx, q, args); err != nil {
		return executeSQLError(err)
	}

	return nil
}

func (r *repository) DeleteTarget(ctx context.Context, targetID string) error {
	sql := `
		DELETE FROM targets WHERE id = $1
	`

	q := database.Query{
		Name: "target_repository.DeleteTarget",
		Sql:  sql,
	}
	if _, err := r.db.DB().ExecContext(ctx, q, targetID); err != nil {
		return executeSQLError(err)
	}

	return nil
}

func (r *repository) CreateTarget(ctx context.Context, missionID string, in dto.CreateTargetParams) (string, error) {
	sql := `
		INSERT INTO targets
		(
			mission_id,
			name,
			country
        )
		VALUES($1,$2,$3)
		RETURNING id
	`
	args := []any{
		missionID,
		in.Name,
		in.Country,
	}

	q := database.Query{
		Name: "target_repository.CreateTarget",
		Sql:  sql,
	}

	var targetID string
	if err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&targetID); err != nil {
		return "", scanRowError(err)
	}

	return targetID, nil
}

func (r *repository) GetTarget(ctx context.Context, targetID string) (entity.Target, error) {
	sql := `
		SELECT id,mission_id,name,country,notes,completed,created_at,updated_at
		FROM targets WHERE id=$1
	`

	q := database.Query{
		Name: "target_repository.GetTarget",
		Sql:  sql,
	}

	var target Target
	row, err := r.db.DB().QueryContext(ctx, q, targetID)
	if err != nil {
		return entity.Target{}, nil
	}
	defer row.Close()

	if err := pgxscan.ScanOne(&target, row); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Target{}, apperrors.TargetNotFoundWithID(targetID)
		}

		return entity.Target{}, scanRowError(err)
	}

	return target.ToServiceEntity(), nil
}

func (r *repository) CountIncompliteTargets(ctx context.Context, missionID string) (int, error) {
	sql := "SELECT id FROM targets WHERE mission_id=$1 AND completed=false;"

	q := database.Query{
		Name: "target_repository.CountIncompliteTargets",
		Sql:  sql,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, missionID)
	if err != nil {
		return 0, executeSQLError(err)
	}
	defer rows.Close()

	count := make([]string, 0, 3)
	if err := pgxscan.ScanAll(&count, rows); err != nil {
		return 0, scanRowsError(err)
	}

	return len(count), nil
}

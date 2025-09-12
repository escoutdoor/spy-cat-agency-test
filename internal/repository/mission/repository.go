package mission

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

var _ def.MissionRepository = (*repository)(nil)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

func New(db database.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetMission(ctx context.Context, missionID string) (entity.Mission, error) {
	sql := `
		SELECT 
			m.id AS id,
			m.cat_id,
			m.completed AS mission_completed,
			m.created_at AS mission_created_at,
			m.updated_at AS mission_updated_at,
			t.id AS target_id,
			t.mission_id as target_mission_id,
			t.name AS target_name,
			t.country AS target_country,
			t.notes AS target_notes,
			t.completed AS target_completed,
			t.created_at AS target_created_at,
			t.updated_at AS target_updated_at
		FROM missions m
		LEFT JOIN targets t ON t.mission_id = m.id
		WHERE m.id=$1
	`

	q := database.Query{
		Name: "mission_repository.GetMission",
		Sql:  sql,
	}
	row, err := r.db.DB().QueryContext(ctx, q, missionID)
	if err != nil {
		return entity.Mission{}, executeSQLError(err)
	}
	defer row.Close()

	var mission MissionRows
	if err := pgxscan.ScanAll(&mission, row); err != nil {
		return entity.Mission{}, scanRowError(err)
	}
	if len(mission) == 0 {
		return entity.Mission{}, apperrors.MissionNotFoundWithID(missionID)
	}

	return mission.ToServiceEntity(), nil
}

func (r *repository) ListMissions(ctx context.Context, limit, offset int) ([]entity.Mission, error) {
	sql := `
		WITH limited_missions AS (
			SELECT *
			FROM missions
			ORDER BY created_at DESC
			LIMIT $1
			OFFSET $2
		)
		SELECT 
			m.id AS id,
			m.cat_id,
			m.completed AS mission_completed,
			m.created_at AS mission_created_at,
			m.updated_at AS mission_updated_at,
			t.id AS target_id,
			t.mission_id as target_mission_id,
			t.name AS target_name,
			t.country AS target_country,
			t.notes AS target_notes,
			t.completed AS target_completed,
			t.created_at AS target_created_at,
			t.updated_at AS target_updated_at
		FROM limited_missions m
		LEFT JOIN targets t ON t.mission_id = m.id
		ORDER BY m.created_at DESC, t.created_at ASC;
	`

	if limit <= 0 {
		limit = defaultLimit
	}
	if offset <= 0 {
		offset = defaultOffset
	}

	q := database.Query{
		Name: "mission_repository.ListMissions",
		Sql:  sql,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, executeSQLError(err)
	}

	var missions MissionRows
	if err := pgxscan.ScanAll(&missions, rows); err != nil {
		return nil, scanRowsError(err)
	}
	defer rows.Close()

	return missions.ToServiceEntities(), nil
}

func (r *repository) UpdateMission(ctx context.Context, in dto.UpdateMissionParams) error {
	sql := `
		UPDATE missions
		SET
	`

	args := pgx.NamedArgs{}
	var updates []string

	if in.CatID != nil {
		args["cat_id"] = in.CatID
		updates = append(updates, "cat_id=@cat_id")
	}
	if in.Completed != nil {
		args["completed"] = in.Completed
		updates = append(updates, "completed=@completed")
	}

	if len(updates) == 0 {
		return apperrors.NoFieldsToUpdate
	}

	sql += fmt.Sprintf("updated_at=now(), %s WHERE id=@mission_id", strings.Join(updates, ", "))
	args["mission_id"] = in.ID

	q := database.Query{
		Name: "mission_repository.UpdateMission",
		Sql:  sql,
	}
	if _, err := r.db.DB().ExecContext(ctx, q, args); err != nil {
		return executeSQLError(err)
	}

	return nil
}

func (r *repository) DeleteMission(ctx context.Context, missionID string) error {
	sql := `
		DELETE FROM missions WHERE id = $1
	`

	q := database.Query{
		Name: "mission_repository.DeleteMission",
		Sql:  sql,
	}
	if _, err := r.db.DB().ExecContext(ctx, q, missionID); err != nil {
		return executeSQLError(err)
	}

	return nil
}

func (r *repository) CreateMission(ctx context.Context, in dto.CreateMissionParams) (string, error) {
	sql := `
		INSERT INTO 
		missions DEFAULT VALUES
		RETURNING id
	`

	q := database.Query{
		Name: "mission_repository.CreateMission",
		Sql:  sql,
	}

	var missionID string
	if err := r.db.DB().QueryRowContext(ctx, q).Scan(&missionID); err != nil {
		return "", scanRowError(err)
	}

	return missionID, nil
}

func (r *repository) IsCatOnMission(ctx context.Context, catID string) (bool, error) {
	sql := `
		SELECT id
		FROM missions WHERE cat_id=$1
	`

	q := database.Query{
		Name: "mission_repository.IsCatOnMission",
		Sql:  sql,
	}

	var missionID string
	if err := r.db.DB().QueryRowContext(ctx, q, catID).Scan(&missionID); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return false, scanRowError(err)
	}

	return missionID != "", nil
}

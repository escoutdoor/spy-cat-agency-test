package cat

import (
	"context"
	"errors"

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

var _ def.CatRepisitory = (*repository)(nil)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

func New(db database.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetCat(ctx context.Context, catID string) (entity.Cat, error) {
	sql := `
		SELECT id,name,years_of_experience,breed,salary,created_at,updated_at FROM cats WHERE ID = $1
	`

	q := database.Query{
		Name: "cat_repository.GetCat",
		Sql:  sql,
	}
	row, err := r.db.DB().QueryContext(ctx, q, catID)
	if err != nil {
		return entity.Cat{}, executeSQLError(err)
	}

	var cat Cat
	if err := pgxscan.ScanOne(&cat, row); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Cat{}, apperrors.CatNotFoundWithID(catID)
		}

		return entity.Cat{}, scanRowError(err)
	}

	return cat.ToServiceEntity(), nil
}

func (r *repository) ListCats(ctx context.Context, limit, offset int) ([]entity.Cat, error) {
	sql := `
		SELECT id,name,years_of_experience,breed,salary,created_at,updated_at FROM cats LIMIT @limit OFFSET @offset
	`
	args := pgx.NamedArgs{
		"limit":  defaultLimit,
		"offset": defaultOffset,
	}

	if limit > 0 {
		args["limit"] = limit
	}
	if offset > 0 {
		args["offset"] = offset
	}

	q := database.Query{
		Name: "cat_repository.ListCats",
		Sql:  sql,
	}
	rows, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return nil, executeSQLError(err)
	}

	var cats Cats
	if err := pgxscan.ScanAll(&cats, rows); err != nil {
		return nil, scanRowsError(err)
	}

	return cats.ToServiceEntities(), nil
}

func (r *repository) UpdateCat(ctx context.Context, in dto.UpdateCatParams) (entity.Cat, error) {
	sql := `
		UPDATE cats
		SET 
			salary=$1,
			updated_at=now()
		WHERE id=$2
		RETURNING id,name,years_of_experience,breed,salary,created_at,updated_at
	`

	q := database.Query{
		Name: "cat_repository.UpdateCat",
		Sql:  sql,
	}
	row, err := r.db.DB().QueryContext(ctx, q, in.Salary, in.ID)
	if err != nil {
		return entity.Cat{}, executeSQLError(err)
	}

	var cat Cat
	if err := pgxscan.ScanOne(&cat, row); err != nil {
		return entity.Cat{}, scanRowError(err)
	}

	return cat.ToServiceEntity(), nil
}

func (r *repository) DeleteCat(ctx context.Context, catID string) error {
	sql := `
		DELETE FROM cats WHERE id = $1
	`

	q := database.Query{
		Name: "cat_repository.DeleteCat",
		Sql:  sql,
	}
	if _, err := r.db.DB().ExecContext(ctx, q, catID); err != nil {
		return executeSQLError(err)
	}

	return nil
}

func (r *repository) CreateCat(ctx context.Context, in dto.CreateCatParams) (entity.Cat, error) {
	sql := `
		INSERT INTO 
		cats(
			name,
			years_of_experience,
			breed,
			salary
        )
		VALUES($1,$2,$3,$4)
		RETURNING id,name,years_of_experience,breed,salary,created_at,updated_at
	`
	args := []any{
		in.Name,
		in.YearsOfExperience,
		in.Breed,
		in.Salary,
	}

	q := database.Query{
		Name: "cat_repository.CreateCat",
		Sql:  sql,
	}
	row, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return entity.Cat{}, executeSQLError(err)
	}

	var cat Cat
	if err := pgxscan.ScanOne(&cat, row); err != nil {
		return entity.Cat{}, scanRowError(err)
	}

	return cat.ToServiceEntity(), nil
}

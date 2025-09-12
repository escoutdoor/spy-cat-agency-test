package cat

import (
	"time"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

type Cat struct {
	ID                string  `db:"id"`
	Name              string  `db:"name"`
	YearsOfExperience int     `db:"years_of_experience"`
	Breed             string  `db:"breed"`
	Salary            float64 `db:"salary"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (e Cat) ToServiceEntity() entity.Cat {
	return entity.Cat{
		ID:                e.ID,
		Name:              e.Name,
		YearsOfExperience: e.YearsOfExperience,
		Breed:             e.Breed,
		Salary:            e.Salary,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
	}
}

type Cats []Cat

func (e Cats) ToServiceEntities() []entity.Cat {
	list := make([]entity.Cat, 0, len(e))
	for _, c := range e {
		list = append(list, c.ToServiceEntity())
	}

	return list
}

func executeSQLError(err error) error {
	return errwrap.Wrap("execute sql", err)
}

func scanRowError(err error) error {
	return errwrap.Wrap("scan row", err)
}

func scanRowsError(err error) error {
	return errwrap.Wrap("scan rows", err)
}

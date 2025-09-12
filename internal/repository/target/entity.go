package target

import (
	"time"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

type Target struct {
	ID        string `db:"id"`
	MissionID string `db:"mission_id"`
	Name      string `db:"name"`
	Country   string `db:"country"`
	Notes     string `db:"notes"`
	Completed bool   `db:"completed"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (e Target) ToServiceEntity() entity.Target {
	return entity.Target{
		ID:        e.ID,
		MissionID: e.MissionID,
		Name:      e.Name,
		Country:   e.Country,
		Notes:     e.Notes,
		Completed: e.Completed,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func executeSQLError(err error) error {
	return errwrap.Wrap("execute sql", err)
}

func scanRowError(err error) error {
	return errwrap.Wrap("scan row", err)
}

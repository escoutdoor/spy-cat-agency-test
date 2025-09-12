package mission

import (
	"time"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

type MissionRow struct {
	ID               string  `db:"id"`
	CatID            *string `db:"cat_id"`
	MissionCompleted bool    `db:"mission_completed"`

	TargetID        string `db:"target_id"`
	TargetMissionID string `db:"target_mission_id"`
	TargetName      string `db:"target_name"`
	TargetCountry   string `db:"target_country"`
	TargetNotes     string `db:"target_notes"`
	TargetCompleted bool   `db:"target_completed"`

	MissionCreatedAt time.Time `db:"mission_created_at"`
	MissionUpdatedAt time.Time `db:"mission_updated_at"`

	TargetCreatedAt time.Time `db:"target_created_at"`
	TargetUpdatedAt time.Time `db:"target_updated_at"`
}

func (e MissionRows) ToServiceEntity() entity.Mission {
	return e.ToServiceEntities()[0]
}

type MissionRows []MissionRow

func (e MissionRows) ToServiceEntities() []entity.Mission {
	rows := make(map[string]MissionRows)
	for _, mr := range e {
		rows[mr.ID] = append(rows[mr.ID], mr)
	}

	missions := make([]entity.Mission, 0, len(rows))
	for id, tgs := range rows {
		targets := []entity.Target{}
		for _, t := range tgs {
			targets = append(targets, entity.Target{
				ID:        t.TargetID,
				MissionID: t.TargetMissionID,
				Name:      t.TargetName,
				Country:   t.TargetCountry,
				Notes:     t.TargetNotes,
				Completed: t.TargetCompleted,
				CreatedAt: t.TargetCreatedAt,
				UpdatedAt: t.TargetUpdatedAt,
			})
		}

		mission := entity.Mission{
			ID:        id,
			Targets:   targets,
			Completed: tgs[0].MissionCompleted,
			CreatedAt: tgs[0].MissionCreatedAt,
			UpdatedAt: tgs[0].MissionUpdatedAt,
		}
		if tgs[0].CatID != nil {
			mission.CatID = tgs[0].CatID
		}
		missions = append(missions, mission)
	}

	return missions
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

package v1

import (
	"fmt"
	"time"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const (
	catIDParam = "catId"
)

type controller struct {
	catService service.CatService
}

func Register(a *fiber.App, catService service.CatService) {
	ctl := &controller{catService: catService}
	r := a.Group("/v1/cats")

	r.Post("/", ctl.createCat)
	r.Delete("/:catId", ctl.deleteCat)
	r.Get("/:catId", ctl.getCat)
	r.Get("/", ctl.listCats)
	r.Patch("/:catId", ctl.updateCat)
}

func validateCatID(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("id parameter is required")
	}

	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid id parameter, should be uuid")
	}

	return nil
}

type catResponse struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	YearsOfExperience int       `json:"yearsOfExperience"`
	Breed             string    `json:"breed"`
	Salary            float64   `json:"salary"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func catToCatResponse(cat entity.Cat) catResponse {
	return catResponse{
		ID:                cat.ID,
		Name:              cat.Name,
		YearsOfExperience: cat.YearsOfExperience,
		Breed:             cat.Breed,
		Salary:            cat.Salary,
		CreatedAt:         cat.CreatedAt,
		UpdatedAt:         cat.UpdatedAt,
	}
}

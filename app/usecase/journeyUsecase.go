package usecase

import (
	"github.com/diegoahg/journey/app/domain"
)

type JourneyRepository interface {
	Save(*domain.Journey) error
	Update(*domain.Journey) error
	GetQueueing() ([]*domain.Journey, error)
	FindByID(int) (*domain.Journey, error)
	RemoveByID(id int) error
}

type JourneyUsecase struct {
	Repo JourneyRepository
}

// JourneyInput takes incoming JSON payload for writing heart rate
type JourneyInput struct {
	ID     int `json:"id"`
	People int `json:"people"`
}

func (c *JourneyUsecase) AddJourneys(ji JourneyInput) error {
	journey := domain.NewJourney(ji.ID, ji.People, 0)
	if err := c.Repo.Save(journey); err != nil {
		return err
	}
	return nil
}

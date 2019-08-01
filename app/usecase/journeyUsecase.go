package usecase

import (
	"github.com/diegoahg/journey/app/domain"
)

// JourneyRepository is a interactor with a Journey Repository
type JourneyRepository interface {
	Save(*domain.Journey) error
	Update(*domain.Journey) error
	GetQueueing() ([]*domain.Journey, error)
	FindByID(int) (*domain.Journey, error)
	RemoveByID(id int) error
}

// JourneyUsecase is in charge to manage journey Repo
type JourneyUsecase struct {
	Repo JourneyRepository
}

// JourneyInput takes incoming JSON payload for writing heart rate
type JourneyInput struct {
	ID     int `json:"id"`
	People int `json:"people"`
}

// AddJourney create and insert a new Journey
func (c *JourneyUsecase) AddJourney(ji JourneyInput) error {
	journey := domain.NewJourney(ji.ID, ji.People, 0)
	if err := c.Repo.Save(journey); err != nil {
		return err
	}
	return nil
}

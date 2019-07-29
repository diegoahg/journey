package repository

import (
	"sync"

	"github.com/diegoahg/journey/app/domain"
)

type journeyRepository struct {
	mu       *sync.Mutex
	journeys map[int]*domain.Journey
}

func NewJourneyRepository() *journeyRepository {
	return &journeyRepository{
		mu:       &sync.Mutex{},
		journeys: map[int]*domain.Journey{},
	}
}
func (r *journeyRepository) FindAll() ([]*domain.Journey, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	journeys := make([]*domain.Journey, len(r.journeys))
	i := 0
	for _, journey := range r.journeys {
		journeys[i] = domain.NewJourney(journey.ID, journey.People, journey.CarID)
		i++
	}
	return journeys, nil
}

func (r *journeyRepository) GetQueueing() ([]*domain.Journey, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	len := 0
	for _, journey := range r.journeys {
		if journey.CarID == 0 {
			len++
		}
	}
	journeys := make([]*domain.Journey, len)
	i := 0
	for _, journey := range r.journeys {
		if journey.CarID == 0 {
			journeys[i] = domain.NewJourney(journey.ID, journey.People, journey.CarID)
			i++
		}
	}
	return journeys, nil
}

func (r *journeyRepository) RemoveAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.journeys = make(map[int]*domain.Journey)
	return nil
}

func (r *journeyRepository) FindByID(id int) (*domain.Journey, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, journey := range r.journeys {
		if journey.ID == id {
			return domain.NewJourney(journey.ID, journey.People, journey.CarID), nil
		}
	}
	return &domain.Journey{}, nil
}

func (r *journeyRepository) Save(journey *domain.Journey) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.journeys[journey.GetID()] = &domain.Journey{
		ID:     journey.GetID(),
		People: journey.GetPeople(),
		CarID:  journey.GetCarID(),
	}
	return nil
}

func (r *journeyRepository) Update(journey *domain.Journey) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.journeys[journey.GetID()].People = journey.GetPeople()
	r.journeys[journey.GetID()].CarID = journey.GetCarID()
	return nil
}

func (r *journeyRepository) RemoveByID(id int) error {
	for i, journey := range r.journeys {
		if journey.ID == id {
			delete(r.journeys, i)
		}
	}
	return nil
}

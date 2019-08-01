package repository

import (
	"sync"

	"github.com/diegoahg/journey/app/domain"
)

// journeyRepository manage the journeys
type journeyRepository struct {
	mu       *sync.Mutex
	journeys map[int]*domain.Journey
}

// NewJourneyRepository instance new journeyRepository
func NewJourneyRepository() *journeyRepository {
	return &journeyRepository{
		mu:       &sync.Mutex{},
		journeys: map[int]*domain.Journey{},
	}
}

// GetQueueing get all of journeys without cars
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

// RemoveAll delete all of journeys
func (r *journeyRepository) RemoveAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.journeys = make(map[int]*domain.Journey)
	return nil
}

// FindByID get a journey by id
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

// Save insert a Journey in the repository
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

// Update modify a Journey in the repository
func (r *journeyRepository) Update(journey *domain.Journey) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.journeys[journey.GetID()].People = journey.GetPeople()
	r.journeys[journey.GetID()].CarID = journey.GetCarID()
	return nil
}

// RemoveByID delete a journey by id
func (r *journeyRepository) RemoveByID(id int) error {
	for i, journey := range r.journeys {
		if journey.ID == id {
			delete(r.journeys, i)
		}
	}
	return nil
}

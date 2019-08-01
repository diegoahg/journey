package repository

import (
	"sync"

	"github.com/diegoahg/journey/app/domain"
)

// carRepository manage the cars
type carRepository struct {
	mu   *sync.Mutex
	cars map[int]*domain.Car
}

// NewCarRepository instance new carRepository
func NewCarRepository() *carRepository {
	return &carRepository{
		mu:   &sync.Mutex{},
		cars: map[int]*domain.Car{},
	}
}

// RemoveAll get all the cars
func (r *carRepository) RemoveAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cars = make(map[int]*domain.Car)
	return nil
}

// FindByID get car by id
func (r *carRepository) FindByID(id int) (*domain.Car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, car := range r.cars {
		if car.ID == id {
			return domain.NewCar(car.ID, car.Seats, car.Empty), nil
		}
	}
	return &domain.Car{}, nil
}

// GetEmptys get cars with seats emptys
func (r *carRepository) GetEmptys() ([]*domain.Car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	len := 0
	for _, car := range r.cars {
		if car.Empty > 0 {
			len++
		}
	}
	cars := make([]*domain.Car, len)
	i := 0
	for _, car := range r.cars {
		if car.Empty > 0 {
			cars[i] = domain.NewCar(car.ID, car.Seats, car.Empty)
			i++
		}
	}
	return cars, nil
}

// Save insert a car in the repository
func (r *carRepository) Save(car *domain.Car) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cars[car.GetID()] = &domain.Car{
		ID:    car.GetID(),
		Seats: car.GetSeats(),
		Empty: car.GetEmpty(),
	}
	return nil
}

// Save Modify a car in the repository
func (r *carRepository) Update(car *domain.Car) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cars[car.GetID()].Seats = car.GetSeats()
	r.cars[car.GetID()].Empty = car.GetEmpty()
	return nil
}

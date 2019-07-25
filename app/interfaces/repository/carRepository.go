package repository

import (
	"sync"

	"github.com/diegoahg/journey/app/domain"
)

type carRepository struct {
	mu   *sync.Mutex
	cars map[int]*domain.Car
}

func NewCarRepository() *carRepository {
	return &carRepository{
		mu:   &sync.Mutex{},
		cars: map[int]*domain.Car{},
	}
}
func (r *carRepository) FindAll() ([]*domain.Car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	cars := make([]*domain.Car, len(r.cars))
	i := 0
	for _, car := range r.cars {
		cars[i] = domain.NewCar(car.ID, car.Seats, car.Empty)
		i++
	}
	return cars, nil
}

func (r *carRepository) RemoveAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cars = make(map[int]*domain.Car)
	return nil
}

func (r *carRepository) FindByID(id int) (*domain.Car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, car := range r.cars {
		if car.ID == id {
			return domain.NewCar(car.ID, car.Seats, car.Empty), nil
		}
	}
	return nil, nil
}

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

func (r *carRepository) Update(car *domain.Car) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cars[car.GetID()].Seats = car.GetSeats()
	r.cars[car.GetID()].Empty = car.GetEmpty()
	return nil
}

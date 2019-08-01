package usecase

import (
	"github.com/diegoahg/journey/app/domain"
)

// CarRepository is a interactor of Car Repository
type CarRepository interface {
	Save(*domain.Car) error
	RemoveAll() error
	GetEmptys() ([]*domain.Car, error)
	Update(*domain.Car) error
	FindByID(int) (*domain.Car, error)
}

// CarUsecase is in charge to do cars actions
type CarUsecase struct {
	Repo CarRepository
}

// CarInput takes incoming JSON payload for writing heart rate
type CarInput struct {
	ID    int `json:"id"`
	Seats int `json:"seats"`
}

// PutCars remove all car data in the repositroy and put news cars.
func (c *CarUsecase) PutCars(cars []CarInput) error {
	if err := c.Repo.RemoveAll(); err != nil {
		return err
	}
	for _, v := range cars {
		car := domain.NewCar(v.ID, v.Seats, v.Seats) // to intialize always Empty === Seats
		if err := c.Repo.Save(car); err != nil {
			return err
		}
	}
	return nil
}

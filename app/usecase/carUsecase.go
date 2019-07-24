package usecase

import (
	"github.com/diegoahg/journey/app/domain"
)

type CarRepository interface {
	Save(*domain.Car) error
	RemoveAll() error
}

type CarUsecase struct {
	Repo CarRepository
}

// CarInput takes incoming JSON payload for writing heart rate
type CarInput struct {
	ID    int `json:"id"`
	Seats int `json:"seats"`
}

func (c *CarUsecase) PutCars(cars []CarInput) error {
	if err := c.Repo.RemoveAll(); err != nil {
		return err
	}
	for _, v := range cars {
		car := domain.NewCar(v.ID, v.Seats)
		if err := c.Repo.Save(car); err != nil {
			return err
		}
	}
	return nil
}

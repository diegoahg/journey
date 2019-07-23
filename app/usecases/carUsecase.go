package usecases

import (
	"github.com/diegoahg/journey/app/domain"
)

type CarUsecase interface {
	Add() (id, seats int)
}

func NewUserUsecase(repo repository.UserRepository, service *service.UserService) *userUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
	}
}

// carUsecase is in charge tu create anew car.
type carUsecase struct {
	Cars []Car
}

// Add create a new Car with a new ID
func (c *CarUsecase) Add(id, seats int) domain.Car {
	var c domain.Car
	c.ID = id
	c.Seats = seats
	c.Cars = append(c.Cars, c)
	return c
}

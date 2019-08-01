package repository

import (
	"testing"

	"github.com/diegoahg/journey/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestCarRepository(t *testing.T) {

	carRepo := NewCarRepository()

	err := carRepo.Save(domain.NewCar(1, 4, 4))
	assert.Nil(t, err)

	err = carRepo.Save(domain.NewCar(2, 5, 5))
	assert.Nil(t, err)

	err = carRepo.Update(domain.NewCar(1, 4, 5))
	assert.Nil(t, err)

	car, err := carRepo.FindByID(1)
	assert.Equal(t, domain.NewCar(1, 4, 5), car)
	assert.Nil(t, err)

	cars, err := carRepo.GetEmptys()
	assert.Equal(t, 1, cars[0].GetID())
	assert.Equal(t, 4, cars[0].GetSeats())
	assert.Equal(t, 5, cars[0].GetEmpty())
	assert.Equal(t, 2, cars[1].GetID())
	assert.Equal(t, 5, cars[1].GetSeats())
	assert.Equal(t, 5, cars[1].GetEmpty())
	assert.Equal(t, 2, len(cars))
	assert.Nil(t, err)

	err = carRepo.RemoveAll()
	assert.Nil(t, err)

	car, err = carRepo.FindByID(1)
	assert.Equal(t, &domain.Car{}, car)
	assert.Nil(t, err)

	car, err = carRepo.FindByID(1)
	assert.Equal(t, &domain.Car{}, car)
	assert.Nil(t, err)
}

package usecase

import (
	"fmt"
	"testing"

	"github.com/diegoahg/journey/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestPutCarsOk(t *testing.T) {

	var mock MockCarRepository

	car := domain.NewCar(1, 4, 4)

	mock.On("RemoveAll").Return(nil)
	mock.On("Save", car).Return(nil)

	u := CarUsecase{
		Repo: &mock,
	}

	input := []CarInput{
		CarInput{
			ID:    1,
			Seats: 4,
		},
	}
	r := u.PutCars(input)
	assert.Nil(t, r)
	mock.AssertExpectations(t)
}

func TestPutCarsErrorRemoveAll(t *testing.T) {

	var mock MockCarRepository

	mock.On("RemoveAll").Return(fmt.Errorf("any"))

	u := CarUsecase{
		Repo: &mock,
	}

	input := []CarInput{
		CarInput{
			ID:    1,
			Seats: 4,
		},
	}
	r := u.PutCars(input)
	assert.Error(t, r)
	mock.AssertExpectations(t)
}

func TestPutCarsErrorSave(t *testing.T) {

	var mock MockCarRepository

	car := domain.NewCar(1, 4, 4)

	mock.On("RemoveAll").Return(nil)
	mock.On("Save", car).Return(fmt.Errorf("any"))

	u := CarUsecase{
		Repo: &mock,
	}

	input := []CarInput{
		CarInput{
			ID:    1,
			Seats: 4,
		},
	}
	r := u.PutCars(input)
	assert.Error(t, r)
	mock.AssertExpectations(t)
}

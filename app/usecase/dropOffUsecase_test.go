package usecase

import (
	"fmt"
	"testing"

	"github.com/diegoahg/journey/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestDropOffOk(t *testing.T) {
	car := &domain.Car{
		ID:    1,
		Seats: 4,
		Empty: 4,
	}
	var mockCarRepo MockCarRepository
	mockCarRepo.On("FindByID", 1).Return(car, nil)
	mockCarRepo.On("Update", car).Return(nil)

	journey := &domain.Journey{
		ID:     1,
		People: 4,
		CarID:  1,
	}
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 1).Return(journey, nil)
	mockJourneyRepo.On("RemoveByID", 1).Return(nil)

	u := DropOffUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
	}

	input := DropOffInput{
		ID: 1,
	}

	r, err := u.DropOff(input)
	assert.Equal(t, r, 200)
	assert.Nil(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestDropOffErrorFindJoruney(t *testing.T) {

	var mockCarRepo MockCarRepository
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 2).Return(&domain.Journey{}, fmt.Errorf("any"))

	u := DropOffUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
	}

	input := DropOffInput{
		ID: 2,
	}

	r, err := u.DropOff(input)
	assert.Equal(t, r, 0)
	assert.Error(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestDropOffErrorJourneyNotFound(t *testing.T) {

	var mockCarRepo MockCarRepository
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 2).Return(&domain.Journey{}, nil)

	u := DropOffUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
	}

	input := DropOffInput{
		ID: 2,
	}

	r, err := u.DropOff(input)
	assert.Equal(t, r, 404)
	assert.Nil(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestDropOffErrorFindCar(t *testing.T) {

	var mockCarRepo MockCarRepository
	mockCarRepo.On("FindByID", 1).Return(&domain.Car{}, fmt.Errorf("any"))
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 2).Return(&domain.Journey{ID: 2, CarID: 1}, nil)

	u := DropOffUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
	}

	input := DropOffInput{
		ID: 2,
	}

	r, err := u.DropOff(input)
	assert.Equal(t, r, 0)
	assert.Error(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestDropOffErrorUpdateCar(t *testing.T) {

	var mockCarRepo MockCarRepository
	mockCarRepo.On("FindByID", 1).Return(&domain.Car{}, nil)
	mockCarRepo.On("Update", &domain.Car{}).Return(fmt.Errorf("any"))
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 2).Return(&domain.Journey{ID: 2, CarID: 1}, nil)

	u := DropOffUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
	}

	input := DropOffInput{
		ID: 2,
	}

	r, err := u.DropOff(input)
	assert.Equal(t, r, 0)
	assert.Error(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestDropOffErrorRemoveJourney(t *testing.T) {

	var mockCarRepo MockCarRepository
	mockCarRepo.On("FindByID", 1).Return(&domain.Car{ID: 1, Seats: 4}, nil)
	mockCarRepo.On("Update", &domain.Car{ID: 1, Seats: 4, Empty: 4}).Return(nil)
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 2).Return(&domain.Journey{ID: 2, People: 4, CarID: 1}, nil)
	mockJourneyRepo.On("RemoveByID", 2).Return(fmt.Errorf("any"))

	u := DropOffUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
	}

	input := DropOffInput{
		ID: 2,
	}

	r, err := u.DropOff(input)
	assert.Equal(t, r, 0)
	assert.Error(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

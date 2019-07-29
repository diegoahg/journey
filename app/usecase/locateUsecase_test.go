package usecase

import (
	"fmt"
	"testing"

	"github.com/diegoahg/journey/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestLocateOk(t *testing.T) {
	journey := &domain.Journey{
		ID:     1,
		People: 4,
		CarID:  1,
	}
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 1).Return(journey, nil)

	u := LocateUsecase{
		JourneyRepo: &mockJourneyRepo,
	}

	input := LocateInput{
		ID: 1,
	}

	r, err := u.Locate(input)
	assert.Equal(t, r, 200)
	assert.Nil(t, err)
	mockJourneyRepo.AssertExpectations(t)
}

func TestLocateErrorFindJoruney(t *testing.T) {

	var mockCarRepo MockCarRepository
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 2).Return(&domain.Journey{}, fmt.Errorf("any"))

	u := LocateUsecase{
		JourneyRepo: &mockJourneyRepo,
	}

	input := LocateInput{
		ID: 2,
	}

	r, err := u.Locate(input)
	assert.Equal(t, r, 0)
	assert.Error(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestLocateErrorJourneyNotFound(t *testing.T) {

	var mockCarRepo MockCarRepository
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 2).Return(&domain.Journey{}, nil)

	u := LocateUsecase{
		JourneyRepo: &mockJourneyRepo,
	}

	input := LocateInput{
		ID: 2,
	}

	r, err := u.Locate(input)
	assert.Equal(t, r, 404)
	assert.Nil(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestLocateErrorJourneyNotCar(t *testing.T) {

	var mockCarRepo MockCarRepository
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("FindByID", 2).Return(&domain.Journey{ID: 2}, nil)

	u := LocateUsecase{
		JourneyRepo: &mockJourneyRepo,
	}

	input := LocateInput{
		ID: 2,
	}

	r, err := u.Locate(input)
	assert.Equal(t, r, 204)
	assert.Nil(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

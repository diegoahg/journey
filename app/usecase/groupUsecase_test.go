package usecase

import (
	"fmt"
	"testing"

	"github.com/diegoahg/journey/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestGroupOk(t *testing.T) {
	cars := []*domain.Car{
		&domain.Car{
			ID:    1,
			Seats: 4,
			Empty: 4,
		},
		&domain.Car{
			ID:    2,
			Seats: 5,
			Empty: 5,
		},
		&domain.Car{
			ID:    3,
			Seats: 6,
			Empty: 6,
		},
	}
	var mockCarRepo MockCarRepository
	mockCarRepo.On("GetEmptys").Return(cars, nil)
	mockCarRepo.On("Update", cars[0]).Return(nil).Once()
	mockCarRepo.On("Update", cars[2]).Return(nil).Once()

	journeys := []*domain.Journey{
		&domain.Journey{
			ID:     1,
			People: 4,
			CarID:  0,
		},
		&domain.Journey{
			ID:     2,
			People: 6,
			CarID:  0,
		},
	}
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("GetQueueing").Return(journeys, nil)
	mockJourneyRepo.On("Update", journeys[0]).Return(nil).Once()
	mockJourneyRepo.On("Update", journeys[1]).Return(nil).Once()

	u := GroupUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
		IsTest:      true,
	}

	err := u.Assign()
	assert.Nil(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestGroupErrorGetEmpty(t *testing.T) {
	var mockCarRepo MockCarRepository
	mockCarRepo.On("GetEmptys").Return([]*domain.Car{}, fmt.Errorf("any"))

	var mockJourneyRepo MockJourneyRepository

	u := GroupUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
		IsTest:      true,
	}

	err := u.Assign()
	assert.Error(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestGroupGetQueueing(t *testing.T) {
	var mockCarRepo MockCarRepository
	mockCarRepo.On("GetEmptys").Return([]*domain.Car{}, nil)

	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("GetQueueing").Return([]*domain.Journey{}, fmt.Errorf("any"))

	u := GroupUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
		IsTest:      true,
	}

	err := u.Assign()
	assert.Error(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

func TestGroupUpdateCar(t *testing.T) {
	cars := []*domain.Car{
		&domain.Car{
			ID:    1,
			Seats: 4,
			Empty: 4,
		},
		&domain.Car{
			ID:    2,
			Seats: 5,
			Empty: 5,
		},
		&domain.Car{
			ID:    3,
			Seats: 6,
			Empty: 6,
		},
	}
	var mockCarRepo MockCarRepository
	mockCarRepo.On("GetEmptys").Return(cars, nil)
	mockCarRepo.On("Update", cars[0]).Return(fmt.Errorf("any")).Once()

	journeys := []*domain.Journey{
		&domain.Journey{
			ID:     1,
			People: 4,
			CarID:  0,
		},
		&domain.Journey{
			ID:     2,
			People: 6,
			CarID:  0,
		},
	}
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("GetQueueing").Return(journeys, nil)

	u := GroupUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
		IsTest:      true,
	}

	err := u.Assign()
	assert.Error(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}
func TestGroupUpdateJourney(t *testing.T) {
	cars := []*domain.Car{
		&domain.Car{
			ID:    1,
			Seats: 4,
			Empty: 4,
		},
		&domain.Car{
			ID:    2,
			Seats: 5,
			Empty: 5,
		},
		&domain.Car{
			ID:    3,
			Seats: 6,
			Empty: 6,
		},
	}
	var mockCarRepo MockCarRepository
	mockCarRepo.On("GetEmptys").Return(cars, nil)
	mockCarRepo.On("Update", cars[0]).Return(nil).Once()

	journeys := []*domain.Journey{
		&domain.Journey{
			ID:     1,
			People: 4,
			CarID:  0,
		},
		&domain.Journey{
			ID:     2,
			People: 6,
			CarID:  0,
		},
	}
	var mockJourneyRepo MockJourneyRepository
	mockJourneyRepo.On("GetQueueing").Return(journeys, nil)
	mockJourneyRepo.On("Update", journeys[0]).Return(fmt.Errorf("any")).Once()

	u := GroupUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
		IsTest:      true,
	}

	err := u.Assign()
	assert.Error(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

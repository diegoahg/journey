package usecase

import (
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
	cars[1].Empty = 0
	cars[2].Empty = 0
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
	journeys[0].CarID = 1
	journeys[1].CarID = 3
	mockJourneyRepo.On("Update", journeys[0]).Return(nil).Once()
	mockJourneyRepo.On("Update", journeys[1]).Return(nil).Once()

	u := GroupUsecase{
		CarRepo:     &mockCarRepo,
		JourneyRepo: &mockJourneyRepo,
	}

	err := u.Assign()
	assert.Nil(t, err)
	mockCarRepo.AssertExpectations(t)
	mockJourneyRepo.AssertExpectations(t)
}

package usecase

import (
	"fmt"
	"testing"

	"github.com/diegoahg/journey/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestAddJourneyOk(t *testing.T) {

	var mock MockJourneyRepository

	car := domain.NewJourney(1, 4, 0)

	mock.On("Save", car).Return(nil)

	u := JourneyUsecase{
		Repo: &mock,
	}

	input := JourneyInput{
		ID:     1,
		People: 4,
	}
	r := u.AddJourney(input)
	assert.Nil(t, r)
	mock.AssertExpectations(t)
}

func TestAddJourneyError(t *testing.T) {

	var mock MockJourneyRepository

	car := domain.NewJourney(1, 4, 0)

	mock.On("Save", car).Return(fmt.Errorf("any"))

	u := JourneyUsecase{
		Repo: &mock,
	}

	input := JourneyInput{
		ID:     1,
		People: 4,
	}
	r := u.AddJourney(input)
	assert.Error(t, r)
	mock.AssertExpectations(t)
}

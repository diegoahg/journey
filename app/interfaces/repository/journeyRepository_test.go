package repository

import (
	"testing"

	"github.com/diegoahg/journey/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestJourneyRepository(t *testing.T) {

	journeyRepo := NewJourneyRepository()

	err := journeyRepo.Save(domain.NewJourney(1, 4, 0))
	assert.Nil(t, err)

	err = journeyRepo.Save(domain.NewJourney(2, 5, 0))
	assert.Nil(t, err)

	err = journeyRepo.Update(domain.NewJourney(1, 4, 1))
	assert.Nil(t, err)

	journey, err := journeyRepo.FindByID(1)
	assert.Equal(t, domain.NewJourney(1, 4, 1), journey)
	assert.Nil(t, err)

	journeys, err := journeyRepo.GetQueueing()
	assert.Equal(t, 2, journeys[0].GetID())
	assert.Equal(t, 5, journeys[0].GetPeople())
	assert.Equal(t, 0, journeys[0].GetCarID())
	assert.Equal(t, 1, len(journeys))
	assert.Nil(t, err)

	err = journeyRepo.RemoveByID(1)
	assert.Nil(t, err)

	err = journeyRepo.RemoveAll()
	assert.Nil(t, err)

	journey, err = journeyRepo.FindByID(1)
	assert.Equal(t, &domain.Journey{}, journey)
	assert.Nil(t, err)

	journey, err = journeyRepo.FindByID(1)
	assert.Equal(t, &domain.Journey{}, journey)
	assert.Nil(t, err)
}

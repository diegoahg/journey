package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJourney(t *testing.T) {
	c := NewJourney(1, 4, 3)
	var typeOf int

	id := c.GetID()
	assert.IsType(t, typeOf, id)
	assert.Equal(t, id, 1)

	people := c.GetPeople()
	assert.IsType(t, typeOf, people)
	assert.Equal(t, people, 4)

	carID := c.GetCarID()
	assert.IsType(t, typeOf, carID)
	assert.Equal(t, carID, 3)
}

package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCar(t *testing.T) {
	c := NewCar(1, 6, 6)
	var typeOf int

	id := c.GetID()
	assert.IsType(t, typeOf, id)
	assert.Equal(t, id, 1)

	seats := c.GetSeats()
	assert.IsType(t, typeOf, seats)
	assert.Equal(t, seats, 6)

	empty := c.GetEmpty()
	assert.IsType(t, typeOf, empty)
	assert.Equal(t, empty, 6)
}

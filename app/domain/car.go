package domain

// Car represent a car entity in the system
type Car struct {
	ID    int
	Seats int
	Empty int
}

// GetID return ID from Car
func (c *Car) GetID() int {
	return c.ID
}

// GetSeats return Seats from Car
func (c *Car) GetSeats() int {
	return c.Seats
}

// GetEmpty return Empty from Car
func (c *Car) GetEmpty() int {
	return c.Empty
}

// NewCar create a new car entity
func NewCar(id, seats, empty int) *Car {
	return &Car{
		ID:    id,
		Seats: seats,
		Empty: empty,
	}
}

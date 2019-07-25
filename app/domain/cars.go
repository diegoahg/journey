package domain

// Car represent a car entity in the system
type Car struct {
	ID    int
	Seats int
	Empty int
}

func (c *Car) GetID() int {
	return c.ID
}

func (c *Car) GetSeats() int {
	return c.Seats
}

func (c *Car) GetEmpty() int {
	return c.Empty
}

func NewCar(id, seats, empty int) *Car {
	return &Car{
		ID:    id,
		Seats: seats,
		Empty: empty,
	}
}

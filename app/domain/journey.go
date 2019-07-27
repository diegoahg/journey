package domain

// Journey represents a set of gruops that can take a car
type Journey struct {
	ID     int
	People int
	CarID  int
}

// GetID return journey ID
func (j *Journey) GetID() int {
	return j.ID
}

// GetPeople return people
func (j *Journey) GetPeople() int {
	return j.People
}

// GetCarID car ID
func (j *Journey) GetCarID() int {
	return j.CarID
}

// NewJourney create a new journey entity
func NewJourney(id, people, carID int) *Journey {
	return &Journey{
		ID:     id,
		People: people,
		CarID:  carID,
	}
}

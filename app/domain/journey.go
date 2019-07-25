package domain

// Journey represents a set of people that can take a car
type Journey struct {
	ID     int
	People int
	CarID  int
}

func (j *Journey) GetID() int {
	return j.ID
}

func (j *Journey) GetPeople() int {
	return j.People
}

func (j *Journey) GetCarID() int {
	return j.CarID
}

func NewJourney(id, people, carID int) *Journey {
	return &Journey{
		ID:     id,
		People: people,
		CarID:  carID,
	}
}

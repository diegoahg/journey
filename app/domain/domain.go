package domain

import "fmt"

// Car represent a car entity in the system
type Car struct {
	ID    int
	Seats int
	Empty int
}

// Journey represents a set of people that can take a car
type Journey struct {
	ID     int
	People int
}

// QueueJourney represents the people that are waiting for a car
type QueueJourney struct {
	Journeys []Journey
}

// Gruop represents a car with gruop of people
type Gruop struct {
	CarID   int
	Journey Journey
}

// AddJourney add a new journey to the queue
func (qj *QueueJourney) AddJourney(id, people int) {
	var j Journey
	j.ID = id
	j.People = people
	qj.Journeys = append(qj.Journeys, g)
}

// AddGroup initialize a group with a car and journey
func (g *Gruop) AddGroup(carID int, j Journey) {
	g.CarID = carID
	g.Journey = j
}

// GetJourney get the journey for a car
func (qj *QueueJourney) GetJourney(requiered int) (Journey, error) {
	var jAux Journey
	for i, j := range qj.Journeys {
		if j.People <= requiered {
			jAux.ID = j.ID
			jAux.People = j.People
			qj.Journeys = append(qj.Journeys[:i], qj.Journeys[i+1:]...)
			return jAux, nil
		}
	}
	return jAux, fmt.Errorf("NO hay")
}

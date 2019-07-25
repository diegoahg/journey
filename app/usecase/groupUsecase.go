package usecase

import (
	"log"
	"time"
)

type GroupUsecase struct {
	CarRepo     CarRepository
	JourneyRepo JourneyRepository
}

func (c *GroupUsecase) Assign() error {
	for {
		log.Println("EMPEZANDO MI GOROUTINE")

		cars, errC := c.CarRepo.GetEmptys()
		if errC != nil {
			return errC
		}

		journeys, errJ := c.JourneyRepo.GetQueueing()
		if errJ != nil {
			return errJ
		}

		for _, car := range cars {
			for _, journey := range journeys {
				log.Println("CAR: %#v", car)
				log.Println("JOURNEY: %#v", journey)
				if journey.GetPeople() <= car.GetEmpty() {
					log.Println("Yajuu Asigno algo")
					car.Empty = car.GetEmpty() - journey.GetPeople()
					if err := c.CarRepo.Update(car); err != nil {
						return err
					}
					journey.CarID = car.ID
					if err := c.JourneyRepo.Update(journey); err != nil {
						return err
					}
				}
			}
		}
		time.Sleep(2 * time.Second)
	}
}

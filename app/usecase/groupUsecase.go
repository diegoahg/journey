package usecase

import (
	"fmt"
	"time"
)

type GroupUsecase struct {
	CarRepo     CarRepository
	JourneyRepo JourneyRepository
}

func (c *GroupUsecase) Assign() error {
	for {
		cars, errC := c.CarRepo.GetEmptys()
		if errC != nil {
			return errC
		}

		journeys, errJ := c.JourneyRepo.GetQueueing()
		if errJ != nil {
			return errJ
		}

		for _, journey := range journeys {
			for _, car := range cars {
				if journey.GetPeople() <= car.GetEmpty() {
					car.Empty = car.GetEmpty() - journey.GetPeople()
					if err := c.CarRepo.Update(car); err != nil {
						return err
					}
					journey.CarID = car.ID
					if err := c.JourneyRepo.Update(journey); err != nil {
						return err
					}
					fmt.Println("HOLAAA")
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}

package usecase

import (
	"time"
)

type GroupUsecase struct {
	CarRepo     CarRepository
	JourneyRepo JourneyRepository
	IsTest      bool
}

func (g *GroupUsecase) Assign() error {
	for {
		cars, errC := g.CarRepo.GetEmptys()
		if errC != nil {
			return errC
		}

		journeys, errJ := g.JourneyRepo.GetQueueing()
		if errJ != nil {
			return errJ
		}

		for _, journey := range journeys {
			for _, car := range cars {
				if journey.GetPeople() <= car.GetEmpty() && journey.GetCarID() == 0 && car.GetEmpty() > 0 {
					car.Empty = car.GetEmpty() - journey.GetPeople()
					if err := g.CarRepo.Update(car); err != nil {
						return err
					}
					journey.CarID = car.ID
					if err := g.JourneyRepo.Update(journey); err != nil {
						return err
					}
				}
			}
		}
		time.Sleep(5 * time.Second)
		if g.IsTest == true {
			break
		}
	}
	return nil
}

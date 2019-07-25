package usecase

import (
	"github.com/diegoahg/journey/app/domain"
)


type LocateUsecase struct {
	CarRepo     CarRepository
	JourneyRepo JourneyRepository
}

// JourneyInput takes incoming JSON payload for writing heart rate
type LocateInput struct {
	ID     int `json:"id"`
}

func (d *LocateUsecase) Locate(li LocateInput) error {
	journey, err := d.JourneyRepo.FindByID(di.ID);
	if err != nil {
		return err
	}

	if(journey == nil){
		return "not found"
	}

	if(journey.GetCarID() == nil){
		return "not assigned"
	}

	return nil
}

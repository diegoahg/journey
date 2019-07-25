package usecase

import (
	"github.com/diegoahg/journey/app/domain"
)


type DropOffUsecase struct {
	CarRepo     CarRepository
	JourneyRepo JourneyRepository
}

// JourneyInput takes incoming JSON payload for writing heart rate
type DropOffInput struct {
	ID     int `json:"id"`
}

func (d *DropOffUsecase) DropOff(di DropOffInput) error {
	journey, err := d.JourneyRepo.FindByID(di.ID);
	if err != nil {
		return err
	}

	if(journey == nil){
		return "not found"
	} 

	car, err := d.CarRepo.FindByID(journey.GetCarID());
	if err != nil {
		return err
	}

	car.Empty = car.GetEmpty() + journey.GetPeople()
	err:= d.CarRepo.Update(journey.GetCarID());
	if err != nil {
		return err
	}

	if err := d.JourneyRepo.RemoveByID(di.ID); err != nil {
		return err
	}

	//DEVOLVER ESTADOS
	// return  "ok"
	return nil
}

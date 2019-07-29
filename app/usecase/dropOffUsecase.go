package usecase

type DropOffUsecase struct {
	CarRepo     CarRepository
	JourneyRepo JourneyRepository
}

// JourneyInput takes incoming JSON payload for writing heart rate
type DropOffInput struct {
	ID int `schema:"id"`
}

func (d *DropOffUsecase) DropOff(di DropOffInput) (int, error) {
	journey, err := d.JourneyRepo.FindByID(di.ID)
	if err != nil {
		return 0, err
	}

	if journey.GetID() == 0 {
		return 404, nil
	}

	car, err := d.CarRepo.FindByID(journey.GetCarID())
	if err != nil {
		return 0, err
	}

	car.Empty = car.GetEmpty() + journey.GetPeople()
	err = d.CarRepo.Update(car)
	if err != nil {
		return 0, err
	}

	if err := d.JourneyRepo.RemoveByID(di.ID); err != nil {
		return 0, err
	}

	return 200, nil
}

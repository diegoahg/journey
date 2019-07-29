package usecase

type LocateUsecase struct {
	JourneyRepo JourneyRepository
}

// JourneyInput takes incoming JSON payload for writing heart rate
type LocateInput struct {
	ID int `schema:"id"`
}

func (d *LocateUsecase) Locate(li LocateInput) (int, error) {
	journey, err := d.JourneyRepo.FindByID(li.ID)
	if err != nil {
		return 0, err
	}

	if journey.GetID() == 0 {
		return 404, nil
	}

	if journey.GetCarID() == 0 {
		return 204, nil
	}

	return 200, nil
}

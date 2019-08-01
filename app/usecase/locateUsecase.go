package usecase

// LocateUsecase is in charge to manage Journey Repository
type LocateUsecase struct {
	JourneyRepo JourneyRepository
}

// LocateInput takes incoming JSON payload for writing heart rate
type LocateInput struct {
	ID int `schema:"id"`
}

// Locate search a jounery a return status
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

package interfaces

import "sync"

type carRepository struct {
	mu   *sync.Mutex
	cars map[string]*car
}

func NewCarRepository() *carRepository {
	return &carRepository{
		mu:   &sync.Mutex{},
		cars: map[string]*car{},
	}
}
func (r *carRepository) FindAll() ([]*model.car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	cars := make([]*model.car, len(r.cars))
	i := 0
	for _, car := range r.cars {
		cars[i] = model.NewCar(car.ID, car.Email)
		i++
	}
	return cars, nil
}
func (r *carRepository) FindByID(id int) (*model.car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, car := range r.cars {
		if car.ID == id {
			return model.NewCar(car.ID, car.Email), nil
		}
	}
	return nil, nil
}
func (r *carRepository) Save(car *model.car) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cars[car.GetID()] = &car{
		ID:    car.GetID(),
		Email: car.GetEmail(),
	}
	return nil
}

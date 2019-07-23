package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diegoahg/journey/app/usecases"
)

type carService struct {
	carUsecase usecases.CarUsecase
}

// CarInput takes incoming JSON payload for writing heart rate
type CarInput struct {
	ID    int    `json:"id"`
	Seats string `json:"seats"`
}

func NewCarService(carUsecase usecase.CarUsecase) *carService {
	return &carService{
		carUsecase: carUsecase,
	}
}

func (c *carService) AddCar(w http.ResponseWriter, r *http.Request) {
	log.Println("CarsService actived")
	contentType := r.Header.Get("Content-type")
	if contentType != "application/json" {
		w.WriteHeader(400)
	}
	var input []CarInput
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		w.WriteHeader(402)
	}

	if err := json.Unmarshal(body, &input); err != nil {
		w.WriteHeader(404)
	}

	_ := c.carUsecase.AddCar()
	log.Println(fmt.Sprintf("Car created"))
	w.WriteHeader(200)
}

package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diegoahg/journey/app/usecase"
)

type carHandler struct {
	CarUsecase *usecase.CarUsecase
}

func NewCarHandler(carUsecase *usecase.CarUsecase) *carHandler {
	return &carHandler{
		CarUsecase: carUsecase,
	}
}

func (c *carHandler) Execute(w http.ResponseWriter, r *http.Request) {
	log.Println("CarsHandler actived")
	contentType := r.Header.Get("Content-type")
	if contentType != "application/json" {
		log.Println(fmt.Errorf("Content Type is not valid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input []usecase.CarInput
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.validate(input); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.CarUsecase.PutCars(input); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(fmt.Sprintf("Car created"))
	w.WriteHeader(http.StatusOK)
	return
}

func (c *carHandler) validate(data *[]usecase.CarInput) error {
	for _, e := range data {
		if e.ID <= 0 || e.Seats <= 0 {
			return fmt.Errorf("Data is not valid")
		}
	}
	return nil
}

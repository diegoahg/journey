package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diegoahg/journey/app/usecase"
)

type journeyHandler struct {
	JourneyUsecase *usecase.JourneyUsecase
}

func NewJourneyHandler(journeyUsecase *usecase.JourneyUsecase) *journeyHandler {
	return &journeyHandler{
		JourneyUsecase: journeyUsecase,
	}
}

func (c *journeyHandler) Execute(w http.ResponseWriter, r *http.Request) {
	log.Println("JourneysHandler actived")
	contentType := r.Header.Get("Content-type")
	if contentType != "application/json" {
		log.Println(fmt.Errorf("Content Type is not valid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input usecase.JourneyInput
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

	if err := c.JourneyUsecase.AddJourneys(input); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(fmt.Sprintf("Journey created"))
	w.WriteHeader(http.StatusOK)
	return
}

func (c *journeyHandler) validate(e usecase.JourneyInput) error {
	if e.ID <= 0 || e.People <= 0 {
		return fmt.Errorf("Data is not valid")
	}
	if e.People > 6 {
		return fmt.Errorf("This quantity of group is not permitted")
	return nil
}

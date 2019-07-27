package interfaces

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diegoahg/journey/app/usecase"
)

type locateHandler struct {
	LocateUsecase *usecase.LocateUsecase
	Schema        SchemaInteractor
}

func NewLocateHandler(locateUsecase *usecase.LocateUsecase, s SchemaInteractor) *locateHandler {
	return &locateHandler{
		LocateUsecase: locateUsecase,
		Schema:        s,
	}
}

func (d *locateHandler) Execute(w http.ResponseWriter, r *http.Request) {
	log.Println("LocatesHandler actived")
	contentType := r.Header.Get("Content-type")
	accept := r.Header.Get("Accept")
	if contentType != "application/x-www-form-urlencoded" || accept != "application/json" {
		log.Println(fmt.Errorf("Content Type is not valid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input usecase.LocateInput

	err = d.Schema.NewDecoder()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = d.Schema.Decode(&input, r.PostForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := d.validate(input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	statusHTTP, err := d.LocateUsecase.Locate(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(fmt.Sprintf("Locate created"))
	w.WriteHeader(statusHTTP)
	return
}

func (c *locateHandler) validate(e usecase.LocateInput) error {
	if e.ID == 0 {
		return fmt.Errorf("Data is not valid")
	}
	return nil
}

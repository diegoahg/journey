package interfaces

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diegoahg/journey/app/usecase"
)

// SchemaInteractor process the form-data request
type SchemaInteractor interface {
	NewDecoder() error
	Decode(interface{}, map[string][]string) error
}

// dropOffHandler manage the drop off usecase

type dropOffHandler struct {
	DropOffUsecase *usecase.DropOffUsecase
	Schema         SchemaInteractor
}

// NewDropOffHandler instance a new dropOffHandler
func NewDropOffHandler(dropOffUsecase *usecase.DropOffUsecase, s SchemaInteractor) *dropOffHandler {
	return &dropOffHandler{
		DropOffUsecase: dropOffUsecase,
		Schema:         s,
	}
}

// Execute process and validate the http request
func (d *dropOffHandler) Execute(w http.ResponseWriter, r *http.Request) {
	log.Println("DropOffsHandler actived")
	contentType := r.Header.Get("Content-type")
	if contentType != "application/x-www-form-urlencoded" {
		log.Println(fmt.Errorf("Content Type is not valid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input usecase.DropOffInput

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

	statusHTTP, err := d.DropOffUsecase.DropOff(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(fmt.Sprintf("DropOff created"))
	w.WriteHeader(statusHTTP)
	return
}

// validate check if the data is correct
func (c *dropOffHandler) validate(e usecase.DropOffInput) error {
	if e.ID == 0 {
		return fmt.Errorf("Data is not valid")
	}
	return nil
}

package app

import (
	"log"

	"github.com/gorilla/mux"
)

// NewRouter create a Mux instance to connect api
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	log.Println("Initialized router")
	return r
}

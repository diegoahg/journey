package main

import (
	"log"
	"net/http"

	infrastructure "github.com/diegoahg/journey/app/infrastructure"
	"github.com/diegoahg/journey/app/interfaces"
	"github.com/diegoahg/journey/app/interfaces/repository"
	"github.com/diegoahg/journey/app/usecase"
)

func main() {
	r := infrastructure.NewRouter()
	s := infrastructure.NewSchema()

	carRepo := repository.NewCarRepository()
	journeyRepo := repository.NewJourneyRepository()

	carHandler := interfaces.NewCarHandler(
		&usecase.CarUsecase{
			Repo: carRepo,
		},
	)

	journeyHandler := interfaces.NewJourneyHandler(
		&usecase.JourneyUsecase{
			Repo: journeyRepo,
		},
	)

	group := usecase.GroupUsecase{
		CarRepo:     carRepo,
		JourneyRepo: journeyRepo,
	}

	dropOffHandler := interfaces.NewDropOffHandler(
		&usecase.DropOffUsecase{
			CarRepo:     carRepo,
			JourneyRepo: journeyRepo,
		},
		s,
	)

	locateHandler := interfaces.NewLocateHandler(
		&usecase.LocateUsecase{
			JourneyRepo: journeyRepo,
		},
		s,
	)

	go group.Assign()

	r.HandleFunc("/status", interfaces.StatusHandler).Methods("GET")
	r.HandleFunc("/cars", carHandler.Execute).Methods("PUT")
	r.HandleFunc("/journey", journeyHandler.Execute).Methods("POST")
	r.HandleFunc("/dropoff", dropOffHandler.Execute).Methods("POST")
	r.HandleFunc("/locate", locateHandler.Execute).Methods("POST")

	log.Println("Api Initialized")
	http.ListenAndServe(":9091", r)

	// go get
	// Documentar
	// TEst en caso de errores form data
	// go test ./... -coverprofile cover.out
	// go tool cover -html=cover.out -o cover.html

}

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

	go group.Assign()

	r.HandleFunc("/status", interfaces.StatusHandler).Methods("GET")
	r.HandleFunc("/cars", carHandler.Execute).Methods("PUT")
	r.HandleFunc("/journey", journeyHandler.Execute).Methods("POST")
	//r.HandleFunc("/dropoff", app.RemoveBasketHandler).Methods("POST")
	//r.HandleFunc("/locate", app.RemoveBasketHandler).Methods("POST")

	log.Println("Api Initialized")
	http.ListenAndServe(":8080", r)

}

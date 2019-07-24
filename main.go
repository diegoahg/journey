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

	carHandler := interfaces.NewCarHandler(
		&usecase.CarUsecase{
			Repo: repository.NewCarRepository(),
		},
	)

	r.HandleFunc("/status", interfaces.StatusHandler).Methods("GET")
	r.HandleFunc("/cars", carHandler.Execute).Methods("PUT")
	//r.HandleFunc("/journey", app.GetAmountHandler).Methods("POST")
	//r.HandleFunc("/dropoff", app.RemoveBasketHandler).Methods("POST")
	//r.HandleFunc("/locate", app.RemoveBasketHandler).Methods("POST")

	log.Println("Api Initialized")
	http.ListenAndServe(":8080", r)

}

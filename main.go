package main

import (
	"log"
	"net/http"

	infrastructure "github.com/diegoahg/journey/app/infrastructure"
	"github.com/diegoahg/journey/app/interfaces"
)

func main() {
	r := infrastructure.NewRouter()

	r.HandleFunc("/status", interfaces.StatusHandler).Methods("GET")
	r.HandleFunc("/cars", interfaces.AddProductHandler).Methods("PUT")
	//r.HandleFunc("/journey", app.GetAmountHandler).Methods("POST")
	//r.HandleFunc("/dropoff", app.RemoveBasketHandler).Methods("POST")
	//r.HandleFunc("/locate", app.RemoveBasketHandler).Methods("POST")

	log.Println("Initialized api")
	http.ListenAndServe(":8080", r)

}

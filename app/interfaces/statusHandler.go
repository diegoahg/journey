package interfaces

import (
	"log"
	"net/http"
)

// StatusHandler process the http request to return if the service is up
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("StatusHandler actived")
	w.WriteHeader(200)
}

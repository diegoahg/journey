package interfaces

import (
	"log"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("StatusHandler actived")
	w.WriteHeader(200)
}

package server

import (
	"log"
	"net/http"

	"github.com/MSLacerda/status-hub/internal/routes"
	"github.com/gorilla/mux"
)

func BuiltServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.StatusHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

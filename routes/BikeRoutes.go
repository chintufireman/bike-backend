package routes

import (
	"bike-website/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func NewBike() {
	r := mux.NewRouter()
	r.HandleFunc("/new-bike", controller.HandleNewBikeData).Methods("POST")

	http.Handle("/", r)

	// Start the HTTP server
	http.ListenAndServe(":9090", nil)
}
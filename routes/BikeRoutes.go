package routes

import (
	"bike-website/controller"
	"bike-website/security"

	"github.com/gorilla/mux"

	"net/http"
)

func NewBike() {
	r := mux.NewRouter()
	r.HandleFunc("/new-bike", controller.HandleNewBikeData).Methods("POST")

	handler := security.CorsConfig(r)

	http.Handle("/", handler)

	// Start the HTTP server
	http.ListenAndServe(":9090", nil)
}

func FetchAllBikes(servePath *http.ServeMux)  {
	r := mux.NewRouter()
	r.HandleFunc("/fetch-all", controller.HandleFetchAllBike).Methods("GET")

	handler := security.CorsConfig(r)

	servePath.Handle("/", handler)

	// Start the HTTP server
	http.ListenAndServe(":9090", nil)
}

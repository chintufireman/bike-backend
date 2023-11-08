package routes

import (
	"bike-website/controller"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewBike() {
	r := mux.NewRouter()
	r.HandleFunc("/new-bike", controller.HandleNewBikeData).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST"},
	})

	handler := c.Handler(r)

	http.Handle("/", handler)

	// Start the HTTP server
	http.ListenAndServe(":9090", nil)
}
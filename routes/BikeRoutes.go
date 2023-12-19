package routes

import (
	"bike-website/controller"
	"bike-website/security"
	"fmt"

	"github.com/gorilla/mux"

	"net/http"
)

/*
func NewBike() {
	r := mux.NewRouter()
	r.HandleFunc("/new-bike", controller.HandleNewBikeData).Methods("POST")

	handler := security.CorsConfig(r)

	http.Handle("/", handler)

	// Start the HTTP server
	http.ListenAndServe(":9090", nil)
}

func FetchAllBikes() {
	r := mux.NewRouter()
	r.HandleFunc("/fetch-all", controller.HandleFetchAllBike).Methods("GET")

	handler := security.CorsConfig(r)

	http.Handle("/", handler)

	// Start the HTTP server
	http.ListenAndServe(":9090", nil)
}

func UploadImage() {
	r := mux.NewRouter()
	r.HandleFunc("/new-bike-image", controller.HandleImageData).Methods("POST")
	handler := security.CorsConfig(r)
	http.Handle("/", handler)
	http.ListenAndServe(":9090", nil)
}

func FetchImage() {
	r := mux.NewRouter()
	r.HandleFunc("/fetch-image", controller.HandleFetchImage).Methods("GET")
	handler := security.CorsConfig(r)
	http.Handle("/", handler)
	http.ListenAndServe(":9090", nil)
}*/

/*
LoggingMiddleware is a middleware function
that logs information about incoming requests.
*/
func LoggingMiddleware(next http.Handler) http.Handler {
	// Return a new http.Handler that wraps the original handler.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Print information about the incoming request.
		fmt.Println("Incoming request:", r.Method, r.URL.Path)
		// Call the ServeHTTP method of the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}

func SetupRoutes() {
	r := mux.NewRouter()

	// Use the logging middleware for all routes
	r.Use(LoggingMiddleware)

	// Define your routes here
	r.HandleFunc("/new-bike", controller.HandleNewBikeData).Methods("POST")
	r.HandleFunc("/fetch-all", controller.HandleFetchAllBike).Methods("GET")
	r.HandleFunc("/new-bike-image", controller.HandleImageData).Methods("POST")
	r.HandleFunc("/fetch-image", controller.HandleFetchImage).Methods("GET")
	r.HandleFunc("/fetch-bike", controller.HandleFetchOneBikeDetail).Methods("GET")
	// Apply CORS configuration
	handler := security.CorsConfig(r)

	// Start the HTTP server
	http.Handle("/", handler)
	http.ListenAndServe(":9090", nil)
}

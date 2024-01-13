package security

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func CorsConfig(router *mux.Router) (h http.Handler) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://172.105.36.88/"},
		AllowedMethods: []string{"GET", "POST"},
	})

	handler := c.Handler(router)
	return handler
}

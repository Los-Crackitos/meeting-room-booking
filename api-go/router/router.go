package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// CreateRouter ...
func CreateRouter() {
	router := mux.NewRouter()

	createUserRouter(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		},
	})

	handler := c.Handler(router)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Print(err)
	}

}

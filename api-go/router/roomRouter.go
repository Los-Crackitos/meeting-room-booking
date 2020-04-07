package router

import (
	"api-go/handlers"

	"github.com/gorilla/mux"
)

func createRoomRouter(router *mux.Router) {
	roomRouter := router.PathPrefix("/room").Subrouter()

	roomRouter.
		HandleFunc("", handlers.CreateRoom).
		Methods("POST")
	roomRouter.
		HandleFunc("", handlers.GetAllRoom).
		Methods("GET")
	roomRouter.
		HandleFunc("/{id}", handlers.GetRoom).
		Methods("GET")
	roomRouter.
		HandleFunc("", handlers.UpdateRoom).
		Methods("PUT")
	roomRouter.
		HandleFunc("/{id}", handlers.DeleteRoom).
		Methods("DELETE")
}

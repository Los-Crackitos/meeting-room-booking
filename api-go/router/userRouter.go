package router

import (
	"api-go/handlers"

	"github.com/gorilla/mux"
)

func createUserRouter(router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("", handlers.CreateUser).Methods("POST")
	userRouter.HandleFunc("", handlers.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.DeleteUser).Methods("DELETE")

}

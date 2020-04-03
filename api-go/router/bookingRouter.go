package router

import (
	"api-go/handlers"

	"github.com/gorilla/mux"
)

func createBookingRouter(router *mux.Router) {
	bookingRouter := router.PathPrefix("/booking").Subrouter()

	bookingRouter.
		HandleFunc("", handlers.CreateBooking).
		Methods("POST")
	bookingRouter.
		HandleFunc("", handlers.GetAllBookings).
		Methods("GET")
	bookingRouter.
		HandleFunc("/{id}", handlers.GetBooking).
		Methods("GET")
	bookingRouter.
		HandleFunc("", handlers.UpdateBooking).
		Methods("PUT")
	bookingRouter.
		HandleFunc("/{id}", handlers.DeleteBooking).
		Methods("DELETE")
}

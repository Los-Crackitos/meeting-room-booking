package handlers

import (
	"api-go/database/dbhandlers"
	"api-go/database/dbmodels"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var booking dbmodels.Booking

	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	if err := dbhandlers.CreateBooking(&booking); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Booking creation failed")

		return
	}

	json.NewEncoder(w).Encode("Booking creation succeed")

}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var booking dbmodels.Booking

	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	if err := dbhandlers.UpdateBooking(&booking); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Booking update failed")

		return
	}

	json.NewEncoder(w).Encode("Booking update succeed")
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var booking dbmodels.Booking

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	booking.ID = id

	if err := dbhandlers.DeleteBooking(&booking); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Booking %d do not exist", id),
		)

		return
	}

	json.NewEncoder(w).Encode("Booking delete succeed")
}

func GetBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var booking dbmodels.Booking

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	booking.ID = id

	if err := dbhandlers.GetBooking(&booking); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Booking %d do not exist", id),
		)

		return
	}
	json.NewEncoder(w).Encode(booking)
}

func GetAllBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	bookings, err := dbhandlers.GetAllBookings()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("An error occured")
		return
	}
	json.NewEncoder(w).Encode(bookings)
}

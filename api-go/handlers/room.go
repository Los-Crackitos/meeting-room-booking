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

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var room dbmodels.Room

	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	if err := dbhandlers.CreateRoom(&room); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Room creation failed")

		return
	}

	json.NewEncoder(w).Encode("Room creation succeed")
}

func GetAllRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rooms, err := dbhandlers.GetAllRooms()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("An error occured")
		return
	}
	json.NewEncoder(w).Encode(rooms)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var room dbmodels.Room

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	room.ID = id

	if err := dbhandlers.GetRoom(&room); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Room %d do not exist", id),
		)

		return
	}
	json.NewEncoder(w).Encode(room)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var room dbmodels.Room

	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	if err := dbhandlers.UpdateRoom(&room); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Room creation failed")

		return
	}

	json.NewEncoder(w).Encode("Room update succeed")
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var room dbmodels.Room
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	room.ID = id

	if err := dbhandlers.DeleteRoom(&room); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Room %d do not exist", id),
		)
		return
	}

	json.NewEncoder(w).Encode("Room delete succeed")
}

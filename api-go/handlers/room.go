package handlers

import (
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"
	"api-go/models"
	"api-go/database/dbhandlers"
	"api-go/database/dbmodels"
	"github.com/gorilla/mux"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var body models.Room

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	if err := dbhandlers.CreateRoom(&body.Params); err != nil {
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

	var param dbmodels.Room

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	param.ID = id
	
	if err := dbhandlers.GetRoom(&param); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Room %d do not exist", id),
		)

		return
	}
	json.NewEncoder(w).Encode(param)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body models.Room

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	if err := dbhandlers.UpdateRoom(&body.Params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Room creation failed")

		return
	}

	json.NewEncoder(w).Encode("Room update succeed")
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var param dbmodels.Room

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	param.ID = id

	if err := dbhandlers.DeleteRoom(&param); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Room %d do not exist", id),
		)

		return
	}

	json.NewEncoder(w).Encode("Room delete succeed")
}

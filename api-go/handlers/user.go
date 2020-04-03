package handlers

import (
	"api-go/database/dbhandlers"
	"api-go/database/dbmodels"
	"api-go/models"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body models.User

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	if err := dbhandlers.CreateUser(&body.Params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User creation failed")

		return
	}

	json.NewEncoder(w).Encode("User creation succeed")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var param dbmodels.User

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	param.ID = id

	if err := dbhandlers.GetUser(&param); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("User %d do not exist", id),
		)

		return
	}
	json.NewEncoder(w).Encode(param)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := dbhandlers.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("An error occured")
		return
	}
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body models.User

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	if err := dbhandlers.UpdateUser(&body.Params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User creation failed")

		return
	}

	json.NewEncoder(w).Encode("User update succeed")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var param dbmodels.User

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	param.ID = id

	if err := dbhandlers.DeleteUser(&param); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("User %d do not exist", id),
		)

		return
	}

	json.NewEncoder(w).Encode("User delete succeed")
}

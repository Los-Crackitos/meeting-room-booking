package handlers

import (
	"api-go/database/dbhandlers"
	"api-go/models"

	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	if err := dbhandlers.CreateUser(user); err != nil {
		// erreur
		w.WriteHeader(http.StatutBadRequest)
		json.NewEncoder(w).Encode("User creation failed")

		return
	}

	json.NewEncoder(w).Encode("User creation succesed")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("ok")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ok")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

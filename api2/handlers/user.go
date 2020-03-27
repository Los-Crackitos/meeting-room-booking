package handlers

import (
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ok")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

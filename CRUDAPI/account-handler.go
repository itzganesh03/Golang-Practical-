package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	Database.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	Database.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	Database.First(&user, mux.Vars(r)["aid"])
	json.NewEncoder(w).Encode(user)
}

func UpdateAccountByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	Database.First(&user, mux.Vars(r)["aid"])
	json.NewDecoder(r.Body).Decode(&user)
	Database.Save(&user)
	json.NewEncoder(w).Encode(user)

}

func DeleteAccountByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	Database.Delete(&user, mux.Vars(r)["aid"])
	json.NewEncoder(w).Encode("Account has been deleted !!")
}

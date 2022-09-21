package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
	r.HandleFunc("/account/{aid}", GetAccountByID).Methods("GET")
	r.HandleFunc("/account", CreateAccount).Methods("POST")
	r.HandleFunc("/account/{aid}", UpdateAccountByID).Methods("PUT")
	r.HandleFunc("/account/{aid}", DeleteAccountByID).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8080", r))
}

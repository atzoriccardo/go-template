package main

import (
	"github.com/gorilla/mux"
	"github.com/riccardoatzori91/go-template/handlers"
)

func setupRouter(router *mux.Router) {
	router.HandleFunc("/customers/{id}", handlers.GetCustomerByID).Methods("GET")
}

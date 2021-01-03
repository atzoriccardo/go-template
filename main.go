package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	setupRouter(router)
	// router.HandleFunc("/customers/{id}", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Header().Set("Content-Type", "application/json")

	// 	params := mux.Vars(r)
	// 	id := params["id"]

	// 	result := Customer{id, "Riccardo", "Atzori"}
	// 	json.NewEncoder(rw).Encode(result)
	// }).Methods("GET")

	http.ListenAndServe(":3000", router)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/riccardoatzori91/go-template/model"
)

//
func GetCustomerByID(rw http.ResponseWriter, r *http.Request) {
	result := model.Customer{"1", "Riccardo", "Atzori"}
	json.NewEncoder(rw).Encode(result)
}

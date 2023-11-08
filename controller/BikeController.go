package controller

import (
	"bike-website/db"
	"bike-website/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleNewBikeData(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request data into a struct
	var data *model.Bike
	data = &model.Bike{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(*data)
	// Process the data
	response, err := db.InsertNewBikeDetails(data)

	// Return a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}



func Demo(w http.ResponseWriter, r *http.Request)  {
	var data *model.Bike
	data = &model.Bike{}

	json.NewDecoder(r.Body).Decode(data)
	response:=data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
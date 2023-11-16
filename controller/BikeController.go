package controller

import (
	"bike-website/db"
	"bike-website/model"
	"bike-website/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HandleNewBikeData(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request data into a struct
	var data *model.Bike
	data = &model.Bike{}
	result, err := utility.MapReqBodyToJson(data, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//type assertion
	data, ok := result.(*model.Bike)
	if !ok {
		return
	}

	response, err := db.InsertNewBikeDetails(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Demo(w http.ResponseWriter, r *http.Request) {
	var data *model.Bike
	data = &model.Bike{}
	result, err := utility.MapReqBodyToJson(data, r)
	if err != nil {
		return
	}
	data, ok := result.(*model.Bike)
	if !ok {
		return
	}

	response := data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func HandleFetchAllBike(w http.ResponseWriter, r *http.Request) {
	var data []*model.Bike
	data, err := db.FetchAllBikes()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := data
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func HandleImageData(w http.ResponseWriter, r *http.Request) {
	// Get the file from the form data
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Image file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file into a byte slice
	byteData, err := ioutil.ReadAll(file)
	if err != nil {

		return
	}

	var imgtmp *model.Image
	imgtmp = &model.Image{
		Data: byteData,
	}
	response, err := db.UploadImageData(imgtmp)
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}
	// Return a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func HandleFetchImage(w http.ResponseWriter, r *http.Request) {
	data, err := db.FetchImage()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Error writing image data", http.StatusInternalServerError)
		return
	}
}

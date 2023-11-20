package controller

import (
	"bike-website/db"
	"bike-website/model"
	"bike-website/utility"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
)

func HandleNewBikeData(w http.ResponseWriter, r *http.Request) {
	//parse multipart form data
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//accessing Json data
	jsonField:=r.FormValue("jsonField")
	fmt.Println(jsonField)
	var bikeData *model.Bike
	bikeData=&model.Bike{}
	error:=json.Unmarshal([]byte(jsonField), bikeData)
	if error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//accessing the file from request
	file, _, fileError := r.FormFile("image")
	if fileError != nil {
		http.Error(w, "Error retrieving the image file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Read the content of the file into a byte array
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading image file", http.StatusInternalServerError)
		return
	}

	// Populate the Image field in the Bike struct
	bikeData.Image = model.Image{Data: imageData}
	// Insert bike details into the database
	response, err := db.InsertNewBikeDetails(bikeData)

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

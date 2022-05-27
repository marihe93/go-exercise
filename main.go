package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func main() {

	//creating a new gorilla mux router
	router := mux.NewRouter()
	//specify endpoints
	router.HandleFunc("/api/encrypt", encryptJSON).Methods("POST")
	router.HandleFunc("/api/decrypt", decryptJSON).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Definining json structure for validation, using go-playground/validator
type jsonVal struct {
	Value string `json:"value" validate:"required"`
}

// Encryption function, using base64
func encryptJSON(w http.ResponseWriter, r *http.Request) {
	var newVal jsonVal
	reqBody, e := ioutil.ReadAll(r.Body)

	if e != nil {
		fmt.Fprintf(w, "Error")
	}

	// Extracting JSON into newVal variable
	json.Unmarshal([]byte(reqBody), &newVal)

	// Using go-playground/validator to validate the JSON format
	validate := validator.New()
	err := validate.Struct(newVal)

	if err != nil {
		// If invalid JSON format returning 400 status code plus an invalid message
		http.Error(w, "Invalid json, please use format \"value\" : \"string\"", 400)
		return
	}

	// Encrypting newVal.Value using Base64
	sEnc := b64.StdEncoding.EncodeToString([]byte(newVal.Value))

	w.Header().Set("Content-Type", "application/json")
	// Returning HTTP status 200
	w.WriteHeader(http.StatusOK)
	// Returning the encoded string
	w.Write([]byte(sEnc))

}

// Decryption function, using base64
func decryptJSON(w http.ResponseWriter, r *http.Request) {

	var newVal jsonVal
	reqBody, e := ioutil.ReadAll(r.Body)

	if e != nil {
		fmt.Fprintf(w, "Error")
	}

	// Extracting JSON into newVal variable
	json.Unmarshal([]byte(reqBody), &newVal)

	// Using go-playground/validator to validate the JSON format
	validate := validator.New()
	err := validate.Struct(newVal)

	if err != nil {
		// If invalid JSON format returning 400 status code plus an invalid message
		http.Error(w, "Invalid json, please use format \"value\" : \"string\"", 400)
		return
	}

	// Decrypting newVal.Value using Base64
	sDec, _ := b64.StdEncoding.DecodeString(newVal.Value)

	w.Header().Set("Content-Type", "application/json")
	// Returning HTTP status 200
	w.WriteHeader(http.StatusOK)
	// Returning the decoded string
	w.Write(sDec)

}

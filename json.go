package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	if statusCode > 499 {
		log.Printf("Server error: %v", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statusCode, errResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	json, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}

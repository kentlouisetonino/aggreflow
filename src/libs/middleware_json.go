package libs

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Returns the JSON encoding.
	data, error := json.Marshal(payload)

	// Error handling.
	if error != nil {
		// Log the failed marshal payload.
		log.Printf("F\033[31m[ERROR]\033[0m Failed to marshal JSON response: %v.", payload)
		// Write the status code in the header.
		w.WriteHeader(500)
		return
	}

	// Add the json content type.
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func responseWithError(w http.ResponseWriter, code int, message string) {
	// Error handling if code is 500 or more.
	if code > 499 {
		log.Println("Responding with 5XX Error: ", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w, code, errorResponse{
		Error: message,
	})
}

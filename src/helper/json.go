package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  // Returns the JSON encoding.
  data, error := json.Marshal(payload)

  // Error handling.
  if error != nil {
    // Log the failed marshal payload.
    log.Printf("Something is wrong. Failed to marshal JSON response: %v.", payload)
    // Write the status code in the header.
    w.WriteHeader(500);
    return;
  }

  // Add the json content type.
  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(data)
}


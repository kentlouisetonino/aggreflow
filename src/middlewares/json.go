package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJSON(writer http.ResponseWriter, code int, payload interface{}) {
  // Returns the JSON encoding.
  data, error := json.Marshal(payload)

  // Error handling.
  if error != nil {
    // Log the failed marshal payload.
    log.Printf("F\033[31m[ERROR]\033[0m Failed to marshal JSON response: %v.", payload)
    // Write the status code in the header.
    writer.WriteHeader(500)
    return
  }

  // Add the json content type.
  writer.Header().Add("Content-Type", "application/json")
  writer.WriteHeader(code)
  writer.Write(data)
}

func ResponseWithError(writer http.ResponseWriter, code int, message string) {
  // Error handling if code is 500 or more.
  if code > 499 {
    log.Println("Responding with 5XX Error: ", message)
  }

  type errorResponse struct {
    Error string `json:"error"`
  }

  ResponseWithJSON(writer, code, errorResponse{
    Error: message,
  })
}


package main

import "net/http"

// This wil intercept if the ready to send the response.
func handler(w http.ResponseWriter, r *http.Request) {
  respondWithJSON(w, 200, struct{}{})
}


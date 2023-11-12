package helper

import "net/http"

// This wil intercept if the api is alive and running.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

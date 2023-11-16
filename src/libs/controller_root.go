package libs

import "net/http"

// This will handle response in returning error..
func handleError(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "Something is wrong. Please try again.")
}

// This wil intercept if the api is alive and running.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}

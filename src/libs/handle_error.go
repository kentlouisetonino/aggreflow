package libs

import "net/http"

// This will handle response in returning error..
func handleError(w http.ResponseWriter, r *http.Request) {
	responWithError(w, 400, "Something is wrong. Please try again.")
}

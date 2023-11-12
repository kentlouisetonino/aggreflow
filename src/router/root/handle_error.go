package root

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	responWithError(w, 400, "Something is wrong. Please try again.")
}

package helper

import "net/http"

func HandleError(w http.ResponseWriter, r *http.Request) {
  responWithError(w, 400, "Something is wrong. Please try again.")
}

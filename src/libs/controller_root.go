package libs

import "net/http"

// This will handle response in returning error..
func handleError(writer http.ResponseWriter, request *http.Request) {
  responseWithError(writer, 400, "Something is wrong. Please try again.")
}

// This wil intercept if the api is alive and running.
func handleHealth(writer http.ResponseWriter, request *http.Request) {
  responseWithJSON(writer, 200, struct{}{})
}


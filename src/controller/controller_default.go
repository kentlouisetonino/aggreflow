package controller

import (
	"log"
	"net/http"

	"github.com/kentlouisetonino/aggreflow/src/helper"
	"github.com/kentlouisetonino/aggreflow/src/middleware"
)

// This will handle response in returning error..
func HandleError(writer http.ResponseWriter, request *http.Request) {
  middleware.ResponseWithError(writer, 400, "Something is wrong. Please try again.")
  log.Printf(helper.Blue + "[api/default/error]" + helper.Reset + " Request received.")
}

// This wil intercept if the api is alive and running.
func HandleHealth(writer http.ResponseWriter, request *http.Request) {
  middleware.ResponseWithJSON(writer, 200, struct{}{})
  log.Printf(helper.Blue + "[api/default/health]" + helper.Reset + " Request received.")
}

package controllers

import (
	"log"
	"net/http"

	"github.com/kentlouisetonino/aggreflow/src/helpers"
	"github.com/kentlouisetonino/aggreflow/src/middlewares"
)

// This will handle response in returning error..
func HandleError(writer http.ResponseWriter, request *http.Request) {
  middlewares.ResponseWithError(writer, 400, "Something is wrong. Please try again.")
  log.Printf(helpers.Blue + "[api/default/error]" + helpers.Reset + " Request received.")
}

// This wil intercept if the api is alive and running.
func HandleHealth(writer http.ResponseWriter, request *http.Request) {
  middlewares.ResponseWithJSON(writer, 200, struct{}{})
  log.Printf(helpers.Blue + "[api/default/health]" + helpers.Reset + " Request received.")
}


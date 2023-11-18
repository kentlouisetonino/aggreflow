package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kentlouisetonino/aggreflow/database/sqlc"
	"github.com/kentlouisetonino/aggreflow/src/helper"
	"github.com/kentlouisetonino/aggreflow/src/middleware"
)

func HandleUserCreate(writer http.ResponseWriter, request *http.Request, apiConfig *helper.DatabaseConfig) {
  type parameters struct {
    Name string
  }

  // Create a decoder instance.
  decoder := json.NewDecoder(request.Body)
  params := parameters{}

  // Decode the parameters.
  decoderError := decoder.Decode(&params)
  if decoderError != nil {
    middleware.ResponseWithError(writer, 400, "Error parsing JSON.")
    log.Fatal(helper.Red + "[api/users/create]" + helper.Reset + " Error decoding the parameters.")
    return
  }

  // Create a user.
  newUser, newUserError := apiConfig.DB.CreateUser(request.Context(), sqlc.CreateUserParams{
    ID:        uuid.New(),
    Name:      params.Name,
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
  })

  // Return an error message.
  if newUserError != nil {
    middleware.ResponseWithError(writer, 400, "Could not create user.")
    log.Fatal(helper.Red + "[api/users/create]" + helper.Reset + " Error creating the user.")
    return
  }

  // Return the successful response.
  middleware.ResponseWithJSON(writer, 200, newUser)
  log.Printf(helper.Blue + "[api/users/create]" + helper.Reset + " Successfully created the user.")
}

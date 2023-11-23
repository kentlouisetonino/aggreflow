package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kentlouisetonino/aggreflow/database/sqlc"
	"github.com/kentlouisetonino/aggreflow/src/helpers"
	"github.com/kentlouisetonino/aggreflow/src/middlewares"
	"github.com/kentlouisetonino/aggreflow/src/models"
)

func HandleUserCreate(writer http.ResponseWriter, request *http.Request, apiConfig *helpers.DatabaseConfig) {
	type parameters struct {
		Name string
	}

	// Create a decoder instance.
	decoder := json.NewDecoder(request.Body)
	params := parameters{}

	// Decode the parameters.
	decoderError := decoder.Decode(&params)
	if decoderError != nil {
		middlewares.ResponseWithError(writer, 400, "Error parsing JSON.")
		log.Fatal(helpers.Red + "[api/users/create]" + helpers.Reset + " Error decoding the parameters.")
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
		middlewares.ResponseWithError(writer, 400, "Could not create user.")
		log.Fatal(helpers.Red + "[api/users/create]" + helpers.Reset + " Error creating the user.")
		return
	}

	// Return the successful response.
	middlewares.ResponseWithJSON(writer, 200, models.DatabaseUserToUser(newUser))
	log.Printf(helpers.Blue + "[api/users/create]" + helpers.Reset + " Successfully created the user.")
}

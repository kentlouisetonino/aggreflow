package libs

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kentlouisetonino/aggreflow/database/sqlc"
	"github.com/kentlouisetonino/aggreflow/src/helper"
)

func handleUserCreate(writer http.ResponseWriter, request *http.Request, apiConfig *helper.DatabaseConfig) {
	type parameters struct {
		Name string
	}

	// Create a decoder instance.
	decoder := json.NewDecoder(request.Body)
	params := parameters{}

	// Decode the parameters.
	decoderError := decoder.Decode(&params)
	if decoderError != nil {
		responWithError(writer, 400, "Error parsing JSON.")
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
		responWithError(writer, 400, "Could not create user.")
		return
	}

	// Return the successful response.
	respondWithJSON(writer, 200, newUser)
}

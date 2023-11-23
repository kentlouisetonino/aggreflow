package routers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kentlouisetonino/aggreflow/src/controllers"
	"github.com/kentlouisetonino/aggreflow/src/helpers"
)

func RouterUsers(apiConfig *helpers.DatabaseConfig) *chi.Mux {
  routerInstance := chi.NewRouter()

  // For users endpoint.
  routerInstance.Post("/create", func(writer http.ResponseWriter, request *http.Request) {
    controllers.HandleUserCreate(writer, request, apiConfig)
  })

  return routerInstance
}


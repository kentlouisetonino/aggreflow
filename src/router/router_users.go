package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kentlouisetonino/aggreflow/src/controller"
	"github.com/kentlouisetonino/aggreflow/src/helper"
)

func RouterUsers(apiConfig *helper.DatabaseConfig) *chi.Mux {
  routerInstance := chi.NewRouter()

  // For users endpoint.
  routerInstance.Post("/create", func(writer http.ResponseWriter, request *http.Request) {
    controller.HandleUserCreate(writer, request, apiConfig)
  })

  return routerInstance
}

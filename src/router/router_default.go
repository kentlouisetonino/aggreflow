package router

import (
	"github.com/go-chi/chi"
	"github.com/kentlouisetonino/aggreflow/src/controller"
	"github.com/kentlouisetonino/aggreflow/src/helper"
)

func RouterDefault(apiConfig *helper.DatabaseConfig) *chi.Mux {
  routerInstance := chi.NewRouter()

  // For default endpoint health and error checker.
  routerInstance.Get("/error", controller.HandleError)
  routerInstance.HandleFunc("/health", controller.HandleHealth)

  return routerInstance
}

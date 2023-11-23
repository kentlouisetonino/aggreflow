package routers

import (
	"github.com/go-chi/chi"
	"github.com/kentlouisetonino/aggreflow/src/controllers"
	"github.com/kentlouisetonino/aggreflow/src/helpers"
)

func RouterDefault(apiConfig *helpers.DatabaseConfig) *chi.Mux {
  routerInstance := chi.NewRouter()

  // For default endpoint health and error checker.
  routerInstance.Get("/error", controllers.HandleError)
  routerInstance.HandleFunc("/health", controllers.HandleHealth)

  return routerInstance
}


package libs

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kentlouisetonino/aggreflow/src/helper"
)

func Router(apiConfig *helper.DatabaseConfig) *chi.Mux {
	routerInstance := chi.NewRouter()

	// For endpoint health and error checker.
	routerInstance.HandleFunc("/health", handleHealth)
	routerInstance.Get("/error", handleError)

	// For users endpoint.
	routerInstance.Post("/users", func(writer http.ResponseWriter, request *http.Request) {
		handleUserCreate(writer, request, apiConfig)
	})

	return routerInstance
}

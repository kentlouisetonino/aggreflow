package root

import "github.com/go-chi/chi"

func Router() *chi.Mux {
	routerInstance := chi.NewRouter()
	routerInstance.HandleFunc("/health", handleHealth)
	routerInstance.Get("/error", handleError)

	return routerInstance
}

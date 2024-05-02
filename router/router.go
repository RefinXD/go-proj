package router

import (
	"github.com/RefinXD/go-proj/controllers"
	"github.com/go-chi/chi/v5"
)
//Initialize the router for all available paths
func InitRouter(handlers []controllers.Handler) *chi.Mux {
	router := chi.NewRouter()
	for _, handler := range handlers{
		handler.SetRouter(router)
	}

	return router
}
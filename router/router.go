package router

import (
	"github.com/RefinXD/go-proj/controllers"
	"github.com/RefinXD/go-proj/service"
	"github.com/go-chi/chi/v5"
)
//Initialize the router for all available paths
func InitRouter() *chi.Mux {
	router := chi.NewRouter()
	empService := service.Instantiate()
    empHandler := controllers.NewEmployeeController(empService)
	empHandler.SetRouter(router)

	return router
}
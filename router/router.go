package router

import (
	"github.com/RefinXD/go-proj/controllers"
	"github.com/RefinXD/go-proj/service"
	"github.com/go-chi/chi/v5"
)

func InitRouter() *chi.Mux {
	api_router := chi.NewRouter()
	//api_router.Get("/employees",controllers.EmployeeController.GetAllHandler)
	empService := service.Instantiate()
    empHandler := controllers.NewEmployeeController(empService)
	empHandler.SetRouter(api_router)

	return api_router
}
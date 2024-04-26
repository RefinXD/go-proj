package router

import (
	// "github.com/RefinXD/go-proj/controllers"
	"github.com/go-chi/chi/v5"
)

func InitRouter() *chi.Mux {
	api_router := chi.NewRouter()
	//api_router.Get("/employees",controllers.EmployeeController.GetAllHandler)

	return api_router
}
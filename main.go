package main

import (
	"fmt"
	"net/http"

	"github.com/RefinXD/go-proj/controllers"
	"github.com/RefinXD/go-proj/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

    empService := service.Instantiate()
    empHandler := controllers.NewEmployeeController(empService)
    
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)


    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))

    })
    empHandler.SetRouter(r)
    fmt.Println("starting")
    http.ListenAndServe(":3000", r)
}

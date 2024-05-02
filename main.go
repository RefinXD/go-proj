package main

import (
	"fmt"
	"net/http"

	"github.com/RefinXD/go-proj/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))

    })
    // init database layer, database := pg.New...
    // init service layer, service.New(database)
    // for proper dependency injection 
    api_router := router.InitRouter() // go convention : no snake_case, router.New(empService)
    r.Mount("/api",api_router)
    fmt.Println("starting")
    http.ListenAndServe(":3000", r)
}

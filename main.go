package main

import (
	"fmt"
	"net/http"

	"github.com/RefinXD/go-proj/controllers"
	"github.com/RefinXD/go-proj/database"
	"github.com/RefinXD/go-proj/database/connection"
	"github.com/RefinXD/go-proj/router"
	"github.com/RefinXD/go-proj/service"
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
    conn,err := connection.ConnectToDB()
    if err != nil{
        fmt.Println(err)
    }
    empService := service.NewEmployeeService(database.New(conn))
    empHanlder := controllers.NewEmployeeHandler(empService)
    handlers := []controllers.Handler{}
    handlers = append(handlers, empHanlder)
    // init database layer, database := pg.New...
    // init service layer, service.New(database)
    // for proper dependency injection 
    apiRouter := router.InitRouter(handlers) // go convention : no snake_case, router.New(empService)
    r.Mount("/api",apiRouter)
    fmt.Println("starting")
    http.ListenAndServe(":3000", r)
}

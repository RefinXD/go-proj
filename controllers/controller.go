package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RefinXD/go-proj/controllers/dto"
	"github.com/RefinXD/go-proj/service"
	"github.com/go-chi/chi/v5"
)

type EmployeeController struct {
	empService service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) EmployeeController {
	return EmployeeController{
		empService: service,
	}
}

func (e EmployeeController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	employees,err := e.empService.GetAllEmployees(context.Background())
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w,"Internal Server Error")
		return;
	}
	empJson, err := json.Marshal(employees)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w,"Internal Server Error")
		return;
	}
	w.Write(empJson)
}

func (e EmployeeController) GetHandler(w http.ResponseWriter, r *http.Request) {
	employees,err := e.empService.GetEmployee(context.Background(),chi.URLParam(r,"id"))
	if err != nil {
		if(err.Error() == "no rows in result set"){
			w.WriteHeader(400)
			fmt.Fprintf(w,"User not found")
		}else{
			w.WriteHeader(500)
			fmt.Println(err)
			fmt.Fprintf(w,"Internal Server Error")
		}
		return
	}
	empJson, err := json.Marshal(employees)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w,"Internal Server Error")
	}
	w.Write(empJson)
}
func (e EmployeeController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var dto dto.EmployeeDTO
	fmt.Println(r.Body)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w,"Bad Request")
	}
	fmt.Println(dto)
	employee,err := e.empService.CreateEmployees(context.Background(), dto)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w,"Internal Server Error")
	}
	empJson, err := json.Marshal(employee)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w,"Internal Server Error")
	}
	w.WriteHeader(200)
	w.Write(empJson)

}

func (e EmployeeController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var dto dto.EmployeeDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w,"Bad Request")
	}
	employee,err := e.empService.UpdateEmployee(context.Background(), dto,chi.URLParam(r,"id"))
	if err != nil {
		if(err.Error() == "no rows in result set"){
			w.WriteHeader(400)
			fmt.Fprintf(w,"User not found")
		}else{
			w.WriteHeader(500)
			fmt.Println(err)
			fmt.Fprintf(w,"Internal Server Error")
		}
		return
	}
	empJson, err := json.Marshal(employee)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w,"Internal Server Error")
		return
	}
	w.WriteHeader(200)
	w.Write(empJson)

}


func (e EmployeeController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	employees,err := e.empService.DeleteEmployee(context.Background(),chi.URLParam(r,"id"))
	if err != nil{
		if(err.Error() == "no rows in result set"){
			w.WriteHeader(400)
			fmt.Fprintf(w,"User not found")
			return
		}else{
			w.WriteHeader(500)
			fmt.Println(err)
			fmt.Fprintf(w,"Internal Server Error")
			return
		}
	}
	empJson, err := json.Marshal(employees)
	if err != nil || empJson != nil {
		w.WriteHeader(500)
		fmt.Println(err)
		fmt.Fprintf(w,"Internal Server Error")
		return
	}
	fmt.Fprintf(w,"Deleted successfully")
}




func (e EmployeeController) SetRouter(r chi.Router) {
	r.Get("/employees", e.GetAllHandler)
	r.Get("/employee/{id}",e.GetHandler)
	r.Post("/employee", e.CreateHandler)
	r.Put("/employee/{id}",e.UpdateHandler)
	r.Delete("/employee/{id}",e.DeleteHandler)
}

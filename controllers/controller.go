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

type EmployeeController interface{
	GetAllHandler(w http.ResponseWriter, r *http.Request)
	GetHandler(w http.ResponseWriter, r *http.Request)
	CreateHandler(w http.ResponseWriter, r *http.Request)
	UpdateHandler(w http.ResponseWriter, r *http.Request) 
	DeleteHandler(w http.ResponseWriter, r *http.Request)

}


type EmployeeControllerImpl struct {
	EmpService service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) EmployeeControllerImpl {
	return EmployeeControllerImpl{
		EmpService: service,
	}
}

func (e EmployeeControllerImpl) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	employees,err := e.EmpService.GetAllEmployees(context.Background())
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

func (e EmployeeControllerImpl) GetHandler(w http.ResponseWriter, r *http.Request) {
	employees,err := e.EmpService.GetEmployee(context.Background(),chi.URLParam(r,"id"))
	if err != nil {
		if(err.Error() == "no rows in result set"){
			w.WriteHeader(404)
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
func (e EmployeeControllerImpl) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var dto dto.EmployeeDTO
	fmt.Println(r.Body)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w,"Bad Request")
	}
	fmt.Println(dto)
	employee,err := e.EmpService.CreateEmployees(context.Background(), dto)
	if err != nil {
		if err.Error() == "duplicate name"{
			w.WriteHeader(400)
			fmt.Fprintf(w,"Bad request:Duplcate name")
			return
		}else{
			w.WriteHeader(500)
			fmt.Println(err.Error())
			fmt.Fprintf(w,"Internal Server Error")
			return
		}
	}
	empJson, err := json.Marshal(employee)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w,"Internal Server Error")
	}
	w.WriteHeader(200)
	w.Write(empJson)

}

func (e EmployeeControllerImpl) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var dto dto.EmployeeDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w,"Bad Request:No ")
	}
	employee,err := e.EmpService.UpdateEmployee(context.Background(), dto,chi.URLParam(r,"id"))
	if err != nil {
		if(err.Error() == "no rows in result set"){
			w.WriteHeader(404)
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


func (e EmployeeControllerImpl) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	employees,err := e.EmpService.DeleteEmployee(context.Background(),chi.URLParam(r,"id"))
	if err != nil{
		if(err.Error() == "no rows in result set"){
			w.WriteHeader(404)
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




func (e EmployeeControllerImpl) SetRouter(r chi.Router) {
	r.Get("/employees", e.GetAllHandler)
	r.Get("/employee/{id}",e.GetHandler)
	r.Post("/employee", e.CreateHandler)
	r.Put("/employee/{id}",e.UpdateHandler)
	r.Delete("/employee/{id}",e.DeleteHandler)
}

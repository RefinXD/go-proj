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

// do we really need this?
type Handler interface{
	GetAllHandler(w http.ResponseWriter, r *http.Request)
	GetHandler(w http.ResponseWriter, r *http.Request)
	CreateHandler(w http.ResponseWriter, r *http.Request)
	UpdateHandler(w http.ResponseWriter, r *http.Request) 
	DeleteHandler(w http.ResponseWriter, r *http.Request)
	SetRouter(r chi.Router)
}


type EmployeeHandler struct {
	// make private
	empService service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) EmployeeHandler {
	return EmployeeHandler{
		empService: service,
	}
}

// what is the difference between (e EmployeeControllerImpl) and (e *EmployeeControllerImpl) --> research this
func (e EmployeeHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	employees,err := e.empService.GetAllEmployees(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // please use constant
		fmt.Fprintf(w,"Internal Server Error")
		return
	}
	empJson, err := json.Marshal(employees)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Internal Server Error")
		return
	}
	w.Write(empJson)
}

func (e EmployeeHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	employees,err := e.empService.GetEmployee(context.Background(),chi.URLParam(r,"id"))
	if err != nil {
		if(err.Error() == "no rows in result set"){	// don't do this
			// 1. for error comparison use errors.Is(err, errTarget)
			// 2. don't introduce dependecy between multiple layers: this is pg specific error
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w,"User not found")
		}else{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			fmt.Fprintf(w,"Internal Server Error")
		}
		return
	}
	empJson, err := json.Marshal(employees)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Internal Server Error")
	}
	w.Write(empJson)
}
func (e EmployeeHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var dto dto.EmployeeDTO	// don't shadow same variable name with package name
	fmt.Println(r.Body)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Bad Request")
	}
	fmt.Println(dto)
	employee,err := e.empService.CreateEmployees(context.Background(), dto)
	if err != nil {
		if err.Error() == "duplicate name"{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w,"Bad request:Duplcate name")
			return
		}else{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			fmt.Fprintf(w,"Internal Server Error")
			return
		}
	}
	empJson, err := json.Marshal(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Internal Server Error")
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(empJson)

}

func (e EmployeeHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var dto dto.EmployeeDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Bad Request:No ")
	}
	employee,err := e.empService.UpdateEmployee(context.Background(), dto,chi.URLParam(r,"id"))
	if err != nil {
		if(err.Error() == "no rows in result set"){
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w,"User not found")
		}else{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			fmt.Fprintf(w,"Internal Server Error")
		}
		return
	}
	empJson, err := json.Marshal(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Internal Server Error")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(empJson)

}


func (e EmployeeHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	employees,err := e.empService.DeleteEmployee(context.Background(),chi.URLParam(r,"id"))
	if err != nil{
		if(err.Error() == "no rows in result set"){
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w,"User not found")
			return
		}else{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			fmt.Fprintf(w,"Internal Server Error")
			return
		}
	}
	empJson, err := json.Marshal(employees)
	if err != nil || empJson != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		fmt.Fprintf(w,"Internal Server Error")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Deleted successfully")
}




func (e EmployeeHandler) SetRouter(r chi.Router) {
	r.Get("/employees", e.GetAllHandler)
	r.Get("/employee/{id}",e.GetHandler)
	r.Post("/employee", e.CreateHandler)
	r.Put("/employee/{id}",e.UpdateHandler)
	r.Delete("/employee/{id}",e.DeleteHandler)
}

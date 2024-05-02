package controllers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RefinXD/go-proj/controllers"
	"github.com/RefinXD/go-proj/controllers/dto"
	"github.com/RefinXD/go-proj/database/models"
	"github.com/stretchr/testify/assert"
)


type MockTestService struct{

}

func (ms MockTestService) GetAllEmployees(ctx context.Context) (emps []models.Employee,err error){
	return nil,nil
}
func (ms MockTestService) GetEmployee(ctx context.Context,id string) (emps *models.Employee,err error) {
	return nil,nil
}
func (ms MockTestService) CreateEmployees(ctx context.Context,data dto.EmployeeDTO) (emp *models.Employee,err error)  {
	return &models.Employee{},nil
}
func (ms MockTestService) UpdateEmployee(ctx context.Context,data dto.EmployeeDTO,id string) (emp *models.Employee,err error) {
	return nil,nil
}
func (ms MockTestService) DeleteEmployee(ctx context.Context,id string) (emps *models.Employee,err error) {
	return nil,nil
}

func TestEmployeeControllerImpl_GetAll(t *testing.T){
	service := new(MockTestService)
	empController := controllers.NewEmployeeHandler(service)
	req, err := http.NewRequest(http.MethodGet, "/api/employees", nil)
	if err != nil {
	t.Error(err)
	}
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(empController.GetAllHandler)
	handler.ServeHTTP(resp,req)
	assert.NoError(t,err)
	assert.Equal(t,http.StatusOK,resp.Code)
}


func TestEmployeeControllerImpl_Get(t *testing.T){
	service := new(MockTestService)
	empController := controllers.NewEmployeeHandler(service)
	req, err := http.NewRequest(http.MethodGet, "/api/employees/3", nil)
	if err != nil {
	t.Error(err)
	}
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(empController.GetHandler)
	handler.ServeHTTP(resp,req)
	assert.NoError(t,err)
	assert.Equal(t,http.StatusOK,resp.Code)
}
func TestEmployeeControllerImpl_Create(t *testing.T){
	service := new(MockTestService)
	empController := controllers.NewEmployeeHandler(service)
	req, err := http.NewRequest(http.MethodGet, "/api/employees", nil)
	if err != nil {
	t.Error(err)
	}
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(empController.CreateHandler)
	handler.ServeHTTP(resp,req)
	assert.NoError(t,err)
	assert.Equal(t,http.StatusOK,resp.Code)
}
func TestEmployeeControllerImpl_Update(t *testing.T){
	service := new(MockTestService)
	empController := controllers.NewEmployeeHandler(service)
	req, err := http.NewRequest(http.MethodGet, "/api/employees", nil)
	if err != nil {
	t.Error(err)
	}
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(empController.UpdateHandler)
	handler.ServeHTTP(resp,req)
	assert.NoError(t,err)
	assert.Equal(t,http.StatusOK,resp.Code)
}

func TestEmployeeControllerImpl_Delete(t *testing.T){
	service := new(MockTestService)
	empController := controllers.NewEmployeeHandler(service)
	req, err := http.NewRequest(http.MethodGet, "/api/employees", nil)
	if err != nil {
	t.Error(err)
	}
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(empController.GetAllHandler)
	handler.ServeHTTP(resp,req)
	assert.NoError(t,err)
	assert.Equal(t,http.StatusOK,resp.Code)
}
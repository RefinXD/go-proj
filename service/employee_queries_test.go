package service_test

import (
	"context"
	"testing"

	"github.com/RefinXD/go-proj/controllers/dto"
	"github.com/RefinXD/go-proj/database/employee_queries"
	"github.com/RefinXD/go-proj/database/models"
	"github.com/RefinXD/go-proj/service"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)



type MockTestDatabase struct{

}



func (db MockTestDatabase) CreateEmployee(ctx context.Context, arg employee_queries.CreateEmployeeParams) (employee_queries.Employee, error){
	return *new(employee_queries.Employee),nil
}
func (db MockTestDatabase) 	DeleteEmployee(ctx context.Context, id int64) error{
	return nil
}
func (db MockTestDatabase) 	GetEmployee(ctx context.Context, id int64) (employee_queries.Employee, error){
	return *new(employee_queries.Employee),nil
}
func (db MockTestDatabase) 	GetEmployeeByName(ctx context.Context, name string) (employee_queries.Employee, error){
	return *new(employee_queries.Employee),nil
}
func (db MockTestDatabase) 	ListEmployees(ctx context.Context) ([]employee_queries.Employee, error){
	return *new([]employee_queries.Employee),nil
}
func (db MockTestDatabase) 	UpdateEmployee(ctx context.Context, arg employee_queries.UpdateEmployeeParams) (employee_queries.Employee, error){
	return *new(employee_queries.Employee),nil
}
func (db MockTestDatabase) 	WithTx(tx pgx.Tx) *employee_queries.Queries{
	return new(employee_queries.Queries)
}


func TestPersonServiceImpl_List(t *testing.T){
	service := new(service.EmployeeServiceImpl)
	service.Db = MockTestDatabase{}
	data,err := service.GetAllEmployees(context.Background())
	assert.NoError(t,err)
	assert.Equal(t,[]models.Employee{},data)
}

func TestPersonServiceImpl_Get(t *testing.T){
	service := new(service.EmployeeServiceImpl)
	service.Db = MockTestDatabase{}
	data,err := service.GetEmployee(context.Background(),"1")
	assert.NoError(t,err) 
	assert.Equal(t,models.Employee{},*data)
}

func TestPersonServiceImpl_Update(t *testing.T){
	service := new(service.EmployeeServiceImpl)
	service.Db = MockTestDatabase{}
	dto := dto.EmployeeDTO{}
	data,err := service.UpdateEmployee(context.Background(),dto,"1")
	assert.NoError(t,err)
	assert.Equal(t,models.Employee{},*data)
}


func TestPersonServiceImpl_Create(t *testing.T){
	service := new(service.EmployeeServiceImpl)
	service.Db = MockTestDatabase{}
	dto := dto.EmployeeDTO{}
	data,err := service.CreateEmployees(context.Background(),dto)
	assert.NoError(t,err)
	assert.Equal(t,models.Employee{},*data)
}

func TestPersonServiceImpl_delete(t *testing.T){
	service := new(service.EmployeeServiceImpl)
	service.Db = MockTestDatabase{}
	data,err := service.DeleteEmployee(context.Background(),"1")
	assert.NoError(t,err)
	assert.Equal(t,models.Employee{},*data)
}
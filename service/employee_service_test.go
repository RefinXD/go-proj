package service_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/RefinXD/go-proj/controllers/dto"
// 	"github.com/RefinXD/go-proj/database"
// 	"github.com/RefinXD/go-proj/database/models"
// 	"github.com/RefinXD/go-proj/service"
// 	"github.com/jackc/pgx/v5"
// 	"github.com/stretchr/testify/assert"
// )



// type MockTestDatabase struct{

// }



// func (db MockTestDatabase) CreateEmployee(ctx context.Context, arg database.CreateEmployeeParams) (database.Employee, error){
// 	return *new(database.Employee),nil
// }
// func (db MockTestDatabase) 	DeleteEmployee(ctx context.Context, id int64) error{
// 	return nil
// }
// func (db MockTestDatabase) 	GetEmployee(ctx context.Context, id int64) (database.Employee, error){
// 	return *new(database.Employee),nil
// }
// func (db MockTestDatabase) 	GetEmployeeByName(ctx context.Context, name string) (database.Employee, error){
// 	return *new(database.Employee),nil
// }
// func (db MockTestDatabase) 	ListEmployees(ctx context.Context) ([]database.Employee, error){
// 	return *new([]database.Employee),nil
// }
// func (db MockTestDatabase) 	UpdateEmployee(ctx context.Context, arg database.UpdateEmployeeParams) (database.Employee, error){
// 	return *new(database.Employee),nil
// }
// func (db MockTestDatabase) 	WithTx(tx pgx.Tx) *database.Queries{
// 	return new(database.Queries)
// }


// func TestPersonServiceImpl_List(t *testing.T){
// 	service := new(service.EmployeeServiceImpl)
// 	service.repo = MockTestDatabase{}
// 	data,err := service.GetAllEmployees(context.Background())
// 	assert.NoError(t,err)
// 	assert.Equal(t,[]models.Employee{},data)
// }

// func TestPersonServiceImpl_Get(t *testing.T){
// 	service := new(service.EmployeeServiceImpl)
// 	service.Db = MockTestDatabase{}
// 	data,err := service.GetEmployee(context.Background(),"1")
// 	assert.NoError(t,err) 
// 	assert.Equal(t,models.Employee{},*data)
// }

// func TestPersonServiceImpl_Update(t *testing.T){
// 	service := new(service.EmployeeServiceImpl)
// 	service.Db = MockTestDatabase{}
// 	dto := dto.EmployeeDTO{}
// 	data,err := service.UpdateEmployee(context.Background(),dto,"1")
// 	assert.NoError(t,err)
// 	assert.Equal(t,models.Employee{},*data)
// }


// func TestPersonServiceImpl_Create(t *testing.T){
// 	service := new(service.EmployeeServiceImpl)
// 	service.Db = MockTestDatabase{}
// 	dto := dto.EmployeeDTO{}
// 	data,err := service.CreateEmployees(context.Background(),dto)
// 	assert.NoError(t,err)
// 	assert.Equal(t,models.Employee{},*data)
// }

// func TestPersonServiceImpl_delete(t *testing.T){
// 	service := new(service.EmployeeServiceImpl)
// 	service.Db = MockTestDatabase{}
// 	data,err := service.DeleteEmployee(context.Background(),"1")
// 	assert.NoError(t,err)
// 	assert.Equal(t,models.Employee{},*data)
// }
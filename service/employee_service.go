package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/RefinXD/go-proj/controllers/dto"
	"github.com/RefinXD/go-proj/database/connection"
	"github.com/RefinXD/go-proj/database/models"
	"github.com/RefinXD/go-proj/utils"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
)

type EmployeeService interface{
	GetAllEmployees(ctx context.Context) (emps []models.Employee,err error)
	GetEmployee(ctx context.Context,id string) (emps *models.Employee,err error) 
	CreateEmployees(ctx context.Context,data dto.EmployeeDTO) (emp *models.Employee,err error) 
	UpdateEmployee(ctx context.Context,data dto.EmployeeDTO,id string) (emp *models.Employee,err error)
	DeleteEmployee(ctx context.Context,id string) (emps *models.Employee,err error)
}



type EmployeeServiceImpl struct{
	repo connection.EmployeeRepository
}

func NewEmployeeService(repository connection.EmployeeRepository) EmployeeServiceImpl {
	return EmployeeServiceImpl{
		repo: repository,
	}
}

//move all of the database related functionalities to the database files
func (e EmployeeServiceImpl) GetAllEmployees(ctx context.Context) ( []models.Employee, error)  {
	employees,err := e.repo.ListEmployees(ctx)
	if err != nil{
		return nil,err
	}
	parsedEmps := utils.ParseEmployees(employees)
	return *parsedEmps,nil
}

func (e EmployeeServiceImpl) GetEmployee(ctx context.Context,id string) (emps *models.Employee,err error)  {
	int_id,err := strconv.Atoi(id)
	if err != nil{
		return nil,err
	}
	employees,err := e.repo.GetEmployee(ctx,int64(int_id))
	if err != nil{
		return nil,err
	}
	parsedEmps := utils.ParseSingleEmployee(employees)
	return parsedEmps,nil
}

func (e EmployeeServiceImpl) CreateEmployees(ctx context.Context,data dto.EmployeeDTO) (emp *models.Employee,err error)  {
	
	validate := validator.New()


	err = validate.Struct(data)
	if err != nil {
		// Validation errors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range validationErrors {
				fmt.Println(err.Field(), err.Tag(), err.Param())
			}
			return nil,err
		} else {
			fmt.Println("Non-validation error:", err)
			return nil,err
	}
}
	employee,err := e.repo.CreateEmployee(ctx,data)
	if err != nil{
		fmt.Println(err.Error())
		var e *pgconn.PgError
		if errors.As(err,&e) && e.Code == "123"{
			return nil,errors.New("duplicate name")
		}
	}
	return utils.ParseSingleEmployee(employee),nil
}


func (e EmployeeServiceImpl) UpdateEmployee(ctx context.Context,data dto.EmployeeDTO,id string) (emp *models.Employee,err error)  {
	// one solution for this is to introduce a 'database' layer on top of sqlc
	int_id,err := strconv.Atoi(id)
	if err != nil{
		return nil,err
	}
	employee,err := e.repo.UpdateEmployee(ctx,data,int64(int_id))
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return utils.ParseSingleEmployee(employee),nil
}

func (e EmployeeServiceImpl) DeleteEmployee(ctx context.Context,id string) (emps *models.Employee,err error)  {
	

	int_id,err := strconv.Atoi(id);
	if err != nil{
		return nil,err
	}
	queried_emp,err := e.repo.GetEmployee(ctx,int64(int_id));
	if err != nil{
		return nil,err
	}
	err = e.repo.Db.DeleteEmployee(ctx,int64(int_id))
	if err != nil{
		return nil,err
	}
	return utils.ParseSingleEmployee(queried_emp),nil
}
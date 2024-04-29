package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/RefinXD/go-proj/controllers/dto"
	"github.com/RefinXD/go-proj/database/connection"
	"github.com/RefinXD/go-proj/database/employee_queries"
	"github.com/RefinXD/go-proj/database/models"
	"github.com/RefinXD/go-proj/utils"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type EmployeeService interface{
	GetAllEmployees(ctx context.Context) (emps []models.Employee,err error)
	GetEmployee(ctx context.Context,id string) (emps *models.Employee,err error) 
	CreateEmployees(ctx context.Context,data dto.EmployeeDTO) (emp *models.Employee,err error) 
	UpdateEmployee(ctx context.Context,data dto.EmployeeDTO,id string) (emp *models.Employee,err error)
	DeleteEmployee(ctx context.Context,id string) (emps *models.Employee,err error)
}



type Database interface{
	CreateEmployee(ctx context.Context, arg employee_queries.CreateEmployeeParams) (employee_queries.Employee, error)
	DeleteEmployee(ctx context.Context, id int64) error
	GetEmployee(ctx context.Context, id int64) (employee_queries.Employee, error)
	ListEmployees(ctx context.Context) ([]employee_queries.Employee, error)
	UpdateEmployee(ctx context.Context, arg employee_queries.UpdateEmployeeParams) (employee_queries.Employee, error)
	WithTx(tx pgx.Tx) *employee_queries.Queries

}


type EmployeeServiceImpl struct{
	Db Database
}

func Instantiate() (service EmployeeServiceImpl) {
	service = *new(EmployeeServiceImpl)
	conn , err := connection.ConnectToDB()
	service.Db = employee_queries.New(conn)
	if err != nil{
		log.Fatal(err)
	}
	return service
}

func (e EmployeeServiceImpl) GetAllEmployees(ctx context.Context) (emps []models.Employee,err error)  {
	
	employees,err := e.Db.ListEmployees(ctx)
	if err != nil{
		return nil,err
	}
	parsedEmps := utils.ParseEmployees(&employees)
	return *parsedEmps,nil
}

func (e EmployeeServiceImpl) GetEmployee(ctx context.Context,id string) (emps *models.Employee,err error)  {
	
	int_id,err := strconv.Atoi(id)
	if err != nil{
		return nil,err
	}
	employees,err := e.Db.GetEmployee(ctx,int64(int_id))
	if err != nil{
		return nil,err
	}
	parsedEmps := utils.ParseSingleEmployee(&employees)
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
	employee,err := e.Db.CreateEmployee(ctx,employee_queries.CreateEmployeeParams{
		Name: data.Name,
		Dob: *utils.TimeToPgDate(&data.Dob),
		Department: data.Department,
		JobTitle: data.JobTitle,
		Address: data.Address,
		JoinDate: *utils.TimeToPgDate(&data.JoinDate),
		
	})
	if err != nil{
		fmt.Println(err.Error())
		var e *pgconn.PgError
		if errors.As(err,&e) && e.Code == "23505"{
			return nil,errors.New("duplicate name")
		}
	}
	return utils.ParseSingleEmployee(&employee),nil
}


func (e EmployeeServiceImpl) UpdateEmployee(ctx context.Context,data dto.EmployeeDTO,id string) (emp *models.Employee,err error)  {
	
	int_id,err := strconv.Atoi(id);
	if err != nil{
		return nil,err
	}
	queried_emp,err := e.Db.GetEmployee(ctx,int64(int_id));
	if err != nil{
		return nil,err
	}
	
	parsed_emp := utils.RemoveEmptyDataToArgs(data,*utils.ParseSingleEmployee(&queried_emp))

	parsed_emp.ID = int64(int_id)
	employee,err := e.Db.UpdateEmployee(ctx,parsed_emp)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return utils.ParseSingleEmployee(&employee),nil
}

func (e EmployeeServiceImpl) DeleteEmployee(ctx context.Context,id string) (emps *models.Employee,err error)  {
	

	int_id,err := strconv.Atoi(id);
	if err != nil{
		return nil,err
	}
	queried_emp,err := e.Db.GetEmployee(ctx,int64(int_id));
	if err != nil{
		return nil,err
	}
	err = e.Db.DeleteEmployee(ctx,int64(int_id))
	if err != nil{
		return nil,err
	}
	return utils.ParseSingleEmployee(&queried_emp),nil
}
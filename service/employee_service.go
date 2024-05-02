package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/RefinXD/go-proj/controllers/dto"
	"github.com/RefinXD/go-proj/database"
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
	CreateEmployee(ctx context.Context, arg database.CreateEmployeeParams) (database.Employee, error)
	DeleteEmployee(ctx context.Context, id int64) error
	GetEmployee(ctx context.Context, id int64) (database.Employee, error)
	ListEmployees(ctx context.Context) ([]database.Employee, error)
	UpdateEmployee(ctx context.Context, arg database.UpdateEmployeeParams) (database.Employee, error)
	WithTx(tx pgx.Tx) *database.Queries

}


type EmployeeServiceImpl struct{
	Db Database
	conn *pgx.Conn
}

func NewEmployeeService(db Database) EmployeeServiceImpl {
	return EmployeeServiceImpl{
		Db: db,
	}
}


func (e EmployeeServiceImpl) GetAllEmployees(ctx context.Context) ( []models.Employee, error)  {
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
	employee,err := e.Db.CreateEmployee(ctx,database.CreateEmployeeParams{
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
	// one solution for this is to introduce a 'database' layer on top of sqlc
	tx,err := e.conn.Begin(ctx)
	if err != nil {
		return nil,err
	}
	int_id,err := strconv.Atoi(id);
	if err != nil{
		return nil,err
	}
	defer tx.Rollback(ctx);
	qtx := e.Db.WithTx(tx)
	queried_emp,err := qtx.GetEmployee(ctx,int64(int_id));
	if err != nil{
		return nil,err
	}
	
	parsed_emp := utils.RemoveEmptyDataToArgs(data,*utils.ParseSingleEmployee(&queried_emp))

	parsed_emp.ID = int64(int_id)
	employee,err := qtx.UpdateEmployee(ctx,parsed_emp)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	tx.Commit(ctx);
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
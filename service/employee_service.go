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
)

type EmployeeService interface {
	GetAllEmployees(ctx context.Context) (emps []models.Employee, err error)
	GetEmployee(ctx context.Context, id string) (emps *models.Employee, err error)
	CreateEmployees(ctx context.Context, data dto.EmployeeDTO) (emp *models.Employee, err error)
	UpdateEmployee(ctx context.Context, data dto.EmployeeDTO, id string) (emp *models.Employee, err error)
	DeleteEmployee(ctx context.Context, id string) (emps *models.Employee, err error)
}

type EmployeeServiceImpl struct {
	repo connection.EmployeeRepository
}

func NewEmployeeService(repository connection.EmployeeRepository) EmployeeServiceImpl {
	return EmployeeServiceImpl{
		repo: repository,
	}
}

// move all of the database related functionalities to the database files
func (e EmployeeServiceImpl) GetAllEmployees(ctx context.Context) ([]models.Employee, error) {
	employees, err := e.repo.ListEmployees(ctx)
	if err != nil {
		return nil, err
	}
	parsedEmps := utils.ParseEmployees(employees)
	return *parsedEmps, nil
}

func (e EmployeeServiceImpl) GetEmployee(ctx context.Context, id string) (emps *models.Employee, err error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	employees, err := e.repo.GetEmployee(ctx, int64(intId))
	if err != nil {
		return nil, err
	}
	parsedEmps := utils.ParseSingleEmployee(employees)
	return parsedEmps, nil
}

func (e EmployeeServiceImpl) CreateEmployees(ctx context.Context, data dto.EmployeeDTO) (emp *models.Employee, err error) {

	validate := validator.New()

	err = validate.Struct(data)
	if err != nil {
		// Validation errors
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, err := range validationErrors {
				fmt.Println(err.Field(), err.Tag(), err.Param())
			}
			return nil, err
		}
	}
	employee, err := e.repo.CreateEmployee(ctx, data)
	if err != nil {

		return nil, err
	}
	return utils.ParseSingleEmployee(employee), nil
}

func (e EmployeeServiceImpl) UpdateEmployee(ctx context.Context, data dto.EmployeeDTO, id string) (emp *models.Employee, err error) {
	// one solution for this is to introduce a 'database' layer on top of sqlc
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	employee, err := e.repo.UpdateEmployee(ctx, data, int64(intId))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return utils.ParseSingleEmployee(employee), nil
}

func (e EmployeeServiceImpl) DeleteEmployee(ctx context.Context, id string) (emps *models.Employee, err error) {

	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	queriedEmp, err := e.repo.GetEmployee(ctx, int64(intId))
	if err != nil {
		return nil, err
	}
	err = e.repo.Db.DeleteEmployee(ctx, int64(intId))
	if err != nil {
		return nil, err
	}
	return utils.ParseSingleEmployee(queriedEmp), nil
}

func (e EmployeeServiceImpl) SetRepo(repo connection.EmployeeRepository) {
	e.repo = repo
}

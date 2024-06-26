package utils

import (
	"fmt"
	"time"

	"github.com/RefinXD/go-proj/controllers/dto"
	"github.com/RefinXD/go-proj/database"
	"github.com/RefinXD/go-proj/database/models"
	"github.com/jackc/pgx/v5/pgtype"
)

// Used to parse database query results into structs to be returned
func ParseEmployees(dbEmps *[]database.Employee) (emps *[]models.Employee) {
	var empArray []models.Employee
	for _, element := range *dbEmps {
		tempEmployee := models.Employee{
			ID:         element.ID,
			Name:       element.Name,
			Department: element.Department,
			JobTitle:   element.JobTitle,
			Address:    element.Address,
			Dob:        element.Dob.Time,
			JoinDate:   element.JoinDate.Time,
		}
		empArray = append(empArray, tempEmployee)
	}
	return &empArray
}

// ParseSingleEmployee Used to parse a single database query result into structs to be returned
func ParseSingleEmployee(dbEmps *database.Employee) (emps *models.Employee) {
	fmt.Println(dbEmps)
	employee := models.Employee{
		ID:         dbEmps.ID,
		Name:       dbEmps.Name,
		Department: dbEmps.Department,
		JobTitle:   dbEmps.JobTitle,
		Address:    dbEmps.Address,
		Dob:        dbEmps.Dob.Time,
		JoinDate:   dbEmps.JoinDate.Time,
	}

	return &employee
}

func ParseEmployeeDTO(dto *dto.EmployeeDTO) *database.CreateEmployeeParams {
	employee := database.CreateEmployeeParams{
		Name:       dto.Name,
		Department: dto.Department,
		JobTitle:   dto.JobTitle,
		Address:    dto.Address,
		Dob:        *TimeToPgDate(&dto.Dob),
		JoinDate:   *TimeToPgDate(&dto.JoinDate),
	}
	return &employee
}

// Converts time.Time to a pgtype.date, used for parsing date from postman
func TimeToPgDate(t *time.Time) *pgtype.Date {
	if t == nil {
		return nil
	}
	year, month, day := t.Date()
	return &pgtype.Date{
		Time:  time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		Valid: true,
	}
}

// This function is to select between the current data in the database(currentData) and the incoming update data.
// If the incoming data is an empty string/default date time, the current data is selected
func RemoveEmptyDataToArgs(data dto.EmployeeDTO, currentData models.Employee) (emp database.UpdateEmployeeParams) {

	emp = database.UpdateEmployeeParams{}
	if data.Name != "" {
		emp.Name = data.Name
	} else {
		emp.Name = currentData.Name
	}
	if data.Address != "" {
		emp.Address = data.Address
	} else {
		emp.Address = currentData.Address
	}
	if data.JobTitle != "" {
		emp.JobTitle = data.JobTitle
	} else {
		emp.JobTitle = currentData.JobTitle
	}
	if data.Department != "" {
		emp.Department = data.Department
	} else {
		emp.Department = currentData.Department
	}
	if data.Dob.String() != "0001-01-01 00:00:00 +0000 UTC" {
		emp.Dob = *TimeToPgDate(&data.Dob)
	} else {
		emp.Dob = *TimeToPgDate(&currentData.Dob)
	}
	if data.JoinDate.String() != "0001-01-01 00:00:00 +0000 UTC" {
		emp.JoinDate = *TimeToPgDate(&data.JoinDate)
	} else {
		emp.JoinDate = *TimeToPgDate(&currentData.JoinDate)
	}
	return emp
}

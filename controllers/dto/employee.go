package dto

import(
	"time"
)

type EmployeeDTO struct{
	Name        string		`validate:"required"`
	Dob         time.Time	`validate:"required"`
	Department  string		`validate:"required"`
	JobTitle    string		`validate:"required"`
	Address     string		`validate:"required"`
	JoinDate    time.Time	`validate:"required"`
}
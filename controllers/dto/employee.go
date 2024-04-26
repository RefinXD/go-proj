package dto

import(
	"time"
)

type EmployeeDTO struct{
	Name        string
	Dob         time.Time
	Department  string
	JobTitle    string
	Address     string
	JoinDate    time.Time
}
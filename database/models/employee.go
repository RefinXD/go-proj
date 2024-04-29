package models

import (
	"time"
)

type Employee struct {
	ID          int64
	Name        string 		
	Dob         time.Time
	Department  string
	JobTitle    string
	Address     string
	JoinDate    time.Time
}

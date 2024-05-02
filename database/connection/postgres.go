package connection


import (
	"context"
	"log"
	"github.com/jackc/pgx/v5"
	"github.com/RefinXD/go-proj/database"
	

)


func ConnectToDB () (*pgx.Conn,error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:password@localhost:5432/postgres")
	if err != nil {
		log.Println(err)
		return nil,err
	}

	return conn,nil
}

type Database interface{
	CreateEmployee(ctx context.Context, arg database.CreateEmployeeParams) (database.Employee, error)
	DeleteEmployee(ctx context.Context, id int64) error
	GetEmployee(ctx context.Context, id int64) (database.Employee, error)
	ListEmployees(ctx context.Context) ([]database.Employee, error)
	UpdateEmployee(ctx context.Context, arg database.UpdateEmployeeParams) (database.Employee, error)
	WithTx(tx pgx.Tx) *database.Queries
}

type EmployeeDatabase struct{
	Db Database
}

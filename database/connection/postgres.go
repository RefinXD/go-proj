package connection

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/RefinXD/go-proj/controllers/dto"
	"github.com/RefinXD/go-proj/database"
	"github.com/RefinXD/go-proj/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)
const NAME_ALREADY_EXIST = "23505"

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

type EmployeeRepository struct{
	Db Database
	Conn pgx.Conn
}

func Instantiate() (EmployeeRepository) {
	// move database init to main
	conn , err := ConnectToDB()
	if err != nil{
		fmt.Println(err)
	}
	employeeDatabase:= database.New(conn)
	
	empDb := EmployeeRepository{
		Db: employeeDatabase,
		Conn: *conn,
	}
	if err != nil{
		log.Fatal(err)
	}
	return empDb
}

func (e EmployeeRepository) CreateEmployee(ctx context.Context, dto dto.EmployeeDTO) (*database.Employee, error){
	args := utils.ParseEmployeeDTO(&dto)
	emp,err := e.Db.CreateEmployee(ctx,*args)
	if err != nil{
		var e *pgconn.PgError
		if errors.As(err,&e) && e.Code == NAME_ALREADY_EXIST{
			return nil,errors.New("duplicate name")
		}
		return nil,err
	}
	return &emp,err
}
func (e EmployeeRepository) DeleteEmployee(ctx context.Context, id int64) error{
	err := e.Db.DeleteEmployee(ctx,id)
	if err != nil{
		if errors.Is(err,pgx.ErrNoRows){
			return errors.New("not found")
		}
		return err
	}
	return nil
}
func (e EmployeeRepository) GetEmployee(ctx context.Context, id int64) (*database.Employee, error){
	emp,err := e.Db.GetEmployee(ctx,id)
	if err != nil{
		if errors.Is(err,pgx.ErrNoRows){
			return nil,errors.New("not found")
		}
		return nil,err
	}
	return &emp,err
}
func (e EmployeeRepository) ListEmployees(ctx context.Context) (*[]database.Employee, error){
	emps,err := e.Db.ListEmployees(ctx)
	if err != nil{
		return nil,err
	}
	return &emps,err
}
func (e EmployeeRepository) UpdateEmployee(ctx context.Context, dto dto.EmployeeDTO,id int64) (*database.Employee, error){
	tx,err := e.BeginTransaction(ctx)
	if err != nil {
		return nil,err
	}
	defer tx.Rollback(ctx);
	qtx := e.WithTx(tx)
	queried_emp,err := qtx.GetEmployee(ctx,id);
	if err != nil{
		if errors.Is(err,pgx.ErrNoRows){
			return nil,errors.New("not found")
		}
		return nil,err
	}
	
	parsed_emp := utils.RemoveEmptyDataToArgs(dto,*utils.ParseSingleEmployee(&queried_emp))

	parsed_emp.ID = id
	employee,err := qtx.UpdateEmployee(ctx,parsed_emp)
	if err != nil{
		return nil,err
	}
	tx.Commit(ctx);

	return &employee,nil
}
func (e EmployeeRepository) WithTx(tx pgx.Tx) *database.Queries{
	return e.Db.WithTx(tx)
}
func (e EmployeeRepository) BeginTransaction(ctx context.Context) (pgx.Tx,error){
	return e.Conn.Begin(ctx)
}

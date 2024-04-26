package connection


import (
	"context"
	"log"
	"github.com/jackc/pgx/v5"

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
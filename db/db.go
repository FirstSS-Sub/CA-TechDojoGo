package mysql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var CTX context.Context

func Initialize() error {
	var err error

	dsn := "mysql:mysql@tcp(db:3306)/ca_tech_dojo_go"
	DB, err = sql.Open("mysql", dsn)
	fmt.Println("connecting to mysql db: ", dsn)

	if err != nil {
		panic(err.Error())
		return err
	}

	CTX = context.Background()
	return nil
}

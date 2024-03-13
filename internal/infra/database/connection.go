package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/claudineyveloso/bookernet.git/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	dbConn, err := sql.Open("postgres", connStr)

	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	//defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	return dbConn, err
}

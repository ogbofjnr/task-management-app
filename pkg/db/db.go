package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ogbofjnr/maze/pkg/config_manager"
	"log"
)

type DB interface {
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	Rebind(query string) string
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	MustExec(query string, args ...interface{}) sql.Result
}

func InitDB() *sqlx.DB {
	appConfig := config_manager.GetConfig("app.conf")
	source := fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConfig.GetString("postgres.host"),
		appConfig.GetInt("postgres.port"),
		appConfig.GetString("postgres.user"),
		appConfig.GetString("postgres.password"),
		appConfig.GetString("postgres.database"),
	)
	conn, err := sqlx.Connect("postgres", source)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v", err)
	}

	err = conn.Ping()

	if err != nil {
		log.Fatalf("couldn't connect to database: %v", err)
	}

	return conn
}

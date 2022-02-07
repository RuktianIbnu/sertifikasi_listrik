package helper

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	// "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// DB ...
type DB struct {
	SQL *sqlx.DB
}

var (
	dbConn = &DB{}
)

// Init ...
func Init() *DB {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db, err := sqlx.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PSWD"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(time.Duration(300 * time.Second))

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	dbConn.SQL = db

	return dbConn
}

func GetConnection() *sqlx.DB {
	if dbConn.SQL == nil {
		Init()
	}
	return dbConn.SQL
}

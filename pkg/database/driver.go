package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func Connect() *sql.DB {
	conn := "postgres://postgres:postgres@localhost:5432/dchat?sslmode=disable"

	connectionStr := os.Getenv("DB_ADDRESS")
	if len(connectionStr) > 0 {
		conn = connectionStr
	}

	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Printf("Fail to openDB: %v \n", err)
	}
	DB = db
	return DB
}

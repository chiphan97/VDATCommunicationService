package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
var DB *sql.DB
func Connect() *sql.DB{
	conn := "postgres://postgres:123456@localhost:15432/dchat?sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Printf("Fail to openDB: %v \n", err)
	}
	DB = db
	return DB
}
package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var conn = `postgres://postgres:postgres@localhost:5432`

func main() {
	connectionStr := os.Getenv("DATABASE_URL")
	if len(connectionStr) > 0 {
		conn = connectionStr
	}
	conn = conn + "?sslmode=disable"

	statement := `SELECT 1 from pg_database WHERE datname='dchat'`
	db, err := sql.Open("postgres", conn)
	rows, err := db.Query(statement)
	if err != nil {
		log.Println(err)
		return
	}

	if rows.Next() {
		//conn := `postgres://postgres:postgres@localhost:5432/dchat?sslmode=disable`
		db, err = sql.Open("postgres", conn)
		if err != nil {
			log.Printf("Fail to openDB: %v \n", err)
			return
		}
	} else {
		statement := `CREATE DATABASE dchat`
		_, err := db.Exec(statement)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(statement)
		//conn = `postgres://postgres:postgres@localhost:5432/dchat?sslmode=disable`
		db, err = sql.Open("postgres", conn)
		if err != nil {
			log.Printf("Fail to openDB: %v \n", err)
			return
		}
	}

	m := Migrator{
		TableName: "migration",
		ctx:       context.Background(),
		Db:        db,
	}
	m.migrate()

	return
}

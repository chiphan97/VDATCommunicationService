package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const DefaultDbConnection = `postgres://postgres:postgres@localhost:5432`
const DefaultDbName = "dchat"

func main() {
	var connectionDbMaster = DefaultDbConnection
	var connectionDb = DefaultDbConnection

	connectionStrEnv := os.Getenv("DATABASE_URL")

	if len(connectionStrEnv) > 0 {
		connectionDbMaster = connectionStrEnv
		connectionDb = connectionStrEnv
	}

	connectionDbMaster += "?sslmode=disable"
	connectionDb += fmt.Sprintf("/%s?sslmode=disable", DefaultDbName)

	statement := `SELECT 1 from pg_database WHERE datname='` + DefaultDbName + `'`
	db, err := sql.Open("postgres", connectionDbMaster)

	if db == nil {
		log.Print("Cannot connect to db")
		return
	}

	rows, err := db.Query(statement)
	if err != nil {
		log.Println(err)
		return
	}

	if rows.Next() {
		db, err = sql.Open("postgres", connectionDb)
		if err != nil {
			log.Printf("Fail to openDB: %v \n", err)
			return
		}
	} else {
		statement := `CREATE DATABASE ` + DefaultDbName
		_, err := db.Exec(statement)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(statement)
		db, err = sql.Open("postgres", connectionDb)
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

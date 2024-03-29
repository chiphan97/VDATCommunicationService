package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	tableName = "migrations"
	location  = "migration/"
)

type Migrator struct {
	TableName  string
	ctx        context.Context
	Db         *sql.DB
	Statements []string
}

func (m Migrator) migrate() error {
	i := 0

	s := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (id serial, name varchar(255),date timestamp default now(),PRIMARY KEY (id));`, m.TableName)
	tx, err := m.Db.BeginTx(m.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	_, execErr := tx.Exec(s)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
		return execErr
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
		return err
	}

	s = fmt.Sprintf(`SELECT name FROM %s order by date desc`, m.TableName)
	rows, err := m.Db.Query(s)
	if err != nil {
		return err
	}
	var name string
	if rows.Next() {
		err = rows.Scan(&name)
	} else {
		name = "0_"
	}
	version := strings.Split(name, "_")
	i, err = strconv.Atoi(version[0])
	if err != nil {
		log.Fatal(err)
		return err
	}
	i = i + 1
	for {
		path := strconv.Itoa(i) + "_description.sql"
		loPath := location + path
		print(Exists(loPath))
		if Exists(loPath) {
			m.Statements, err = ReadLines(loPath)
			if err != nil {
				log.Println(err)
				break
			}
			log.Println(path)
			for _, statement := range m.Statements {
				tx, err := m.Db.BeginTx(m.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
				if err != nil {
					log.Fatal(err)
				}
				_, execErr := tx.Exec(statement)
				if execErr != nil {
					_ = tx.Rollback()
					log.Fatal(execErr)
				}
				if err := tx.Commit(); err != nil {
					log.Fatal(err)
				}
				log.Println(statement)
			}
			s = fmt.Sprintf(`INSERT INTO %s(name) VALUES ($1) `, m.TableName)
			_, err = m.Db.Exec(s, path)
			if err != nil {
				log.Fatal(err)
				return err
			}
			log.Println(s)
			i++
		} else {
			log.Println("File not exist")
			m.Db.Close()
			return nil
		}
	}
	m.Db.Close()
	return nil
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

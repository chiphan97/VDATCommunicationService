package database
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
const (
	port     = "5432"
	user     = "postgres"
	password = "123456"
	dbname   = "dchat"
)

var DB *sql.DB
func Connect() *sql.DB{
	conn := fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=disable", port, user, password, dbname)
	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Printf("Fail to openDB: %v \n", err)
	}
	DB = db
	return DB
}
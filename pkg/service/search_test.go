package service

import (
	"database/sql"
	"testing"
)

func TestSearch(t *testing.T) {
	db, _ := sql.Open("", "")

	_, _ = db.Query("SELECT * FROM ...")

	t.Error("error")
}

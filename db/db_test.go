package db

import (
	"log"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSqlite(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	_ = db
	log.Println(db)
}

package apiv1

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var h *Handler

func TestMain(m *testing.M) {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dbConnStr := fmt.Sprintf("root:root@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	h = &Handler{
		DB: db,
	}
	m.Run()
}

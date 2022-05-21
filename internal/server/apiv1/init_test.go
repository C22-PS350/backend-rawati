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
	dbConnStr := dbConnStrBuilder()
	if dbConnStr == "" {
		os.Exit(1)
	}
	db, err := gorm.Open(mysql.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	h = &Handler{
		DB: db,
	}
	m.Run()
}

func dbConnStrBuilder() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	switch "" {
	case dbHost, dbPort, dbUsername, dbPassword, dbName:
		return ""
	}

	dbConnStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True"
	dbConnStr = fmt.Sprintf(dbConnStr, dbUsername, dbPassword, dbHost, dbPort, dbName)
	return dbConnStr
}

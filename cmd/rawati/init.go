package main

import (
	"fmt"
	"os"
)

var (
	appHost    = os.Getenv("APP_HOST")
	appPort    = os.Getenv("APP_PORT")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")
	dbConnStr  = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True"
)

func init() {
	if appHost == "" {
		appHost = "0.0.0.0"
	}
	if appPort == "" {
		appPort = "8080"
	}
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbUsername == "" {
		dbUsername = "root"
	}
	if dbPassword == "" {
		dbPassword = "root"
	}
	if dbName == "" {
		dbName = "rawati"
	}

	dbConnStr = fmt.Sprintf(dbConnStr, dbUsername, dbPassword, dbHost, dbPort, dbName)
}

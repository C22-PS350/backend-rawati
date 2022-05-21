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
	switch "" {
	case appHost:
		appHost = "0.0.0.0"
		fallthrough
	case appPort:
		appPort = "8080"
		fallthrough
	case dbHost:
		dbHost = "localhost"
		fallthrough
	case dbPort:
		dbPort = "3306"
		fallthrough
	case dbUsername:
		dbUsername = "root"
		fallthrough
	case dbPassword:
		dbPassword = "root"
		fallthrough
	case dbName:
		dbName = "mars"
	}

	dbConnStr = fmt.Sprintf(dbConnStr, dbUsername, dbPassword, dbHost, dbPort, dbName)
}

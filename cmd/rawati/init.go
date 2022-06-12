package main

import (
	"fmt"
	"os"
)

var (
	environment = os.Getenv("ENVIRONMENT")
	appHost     = os.Getenv("APP_HOST")
	appPort     = os.Getenv("APP_PORT")
	gcpProject  = os.Getenv("GCP_PROJECT")
	modelAPIUrl = os.Getenv("MODEL_API_URL")
	dbConnStr   = ""
)

func init() {
	if environment == "" {
		environment = "local-development"
	}
	if appHost == "" {
		appHost = "0.0.0.0"
	}
	if appPort == "" {
		appPort = "8080"
	}

	switch environment {
	case "local-development":
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbName := os.Getenv("DB_NAME")

		if dbHost == "" {
			dbHost = "127.0.0.1"
		}
		if dbPort == "" {
			dbPort = "3306"
		}
		if dbUser == "" {
			dbUser = "root"
		}
		if dbPass == "" {
			dbPass = "root"
		}
		if dbName == "" {
			dbName = "rawati"
		}

		dbConnStr = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta"
		dbConnStr = fmt.Sprintf(dbConnStr, dbUser, dbPass, dbHost, dbPort, dbName)
		modelAPIUrl = "http://localhost:5000/predictions"

	case "remote-development":
		instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbName := os.Getenv("DB_NAME")
		modelAPIUrl = os.Getenv("MODEL_API_URL")

		if instanceConnectionName == "" {
			panic("INSTANCE_CONNECTION_NAME env is not set")
		}
		if dbUser == "" {
			panic("DB_USER env is not set")
		}
		if dbPass == "" {
			panic("DB_PASS env is not set")
		}
		if dbName == "" {
			panic("DB_NAME env is not set")
		}
		if gcpProject == "" {
			panic("GCP_PROJECT env is not set")
		}
		if modelAPIUrl == "" {
			panic("MODEL_API_URL env is not set")
		}

		dbConnStr = "%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta"
		dbConnStr = fmt.Sprintf(dbConnStr, dbUser, dbPass, instanceConnectionName, dbName)

	default:
		panic("ENVIRONMENT env is invalid")
	}
}

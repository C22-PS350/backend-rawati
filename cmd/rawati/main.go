package main

import (
	"log"

	"github.com/C22-PS350/backend-rawati/internal/server"
)

// @title                         Rawati API
// @version                       1.0
// @description                   Still in development. Updated regularly.
// @host                          localhost:8080
// @BasePath                      /api/v1
// @securityDefinitions.apiToken  ApiToken
// @in                            header
// @name                          Authorization
func main() {
	srvcfg := server.Config{
		Environment:  environment,
		AppHost:      appHost,
		AppPort:      appPort,
		DBConnString: dbConnStr,
	}

	srv := server.New(&srvcfg)
	log.Printf("starting server on %s:%s\n", srv.Config.AppHost, srv.Config.AppPort)
	if err := srv.Start(); err != nil {
		log.Panicf("error starting server: %s", err)
	}
}

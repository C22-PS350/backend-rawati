package main

import (
	"fmt"
	"log"

	"github.com/C22-PS350/backend-rawati/internal/server"
)

// @Title                         Rawati API
// @Version                       1.0
// @Description                   Still in development. Updated regularly.
// @Host                          https://rawati-backend-7wglfeuiba-as.a.run.app
// @BasePath                      /api/v1
// @SecurityDefinitions.apiToken  ApiToken
// @In                            header
// @Name                          Authorization
func main() {
	srvcfg := server.Config{
		Environment:  environment,
		AppHost:      appHost,
		AppPort:      appPort,
		DBConnString: dbConnStr,
		GCPProject:   gcpProject,
		ModelAPIUrl:  modelAPIUrl,
	}

	srv := server.New(&srvcfg)
	fmt.Printf("starting server on %s:%s\n", srv.Config.AppHost, srv.Config.AppPort)
	if err := srv.Start(); err != nil {
		log.Panicf("error starting server: %s", err)
	}
}

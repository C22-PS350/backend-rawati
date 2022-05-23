package main

import (
	"log"

	"github.com/C22-PS350/backend-rawati/internal/server"
)

func main() {
	srvcfg := server.Config{
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

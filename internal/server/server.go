package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/server/apiv1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/patrickmn/go-cache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	Config *Config
}

func New(cfg *Config) *Server {
	return &Server{Config: cfg}
}

func (srv *Server) Start() error {
	router, err := srv.createRouter()
	if err != nil {
		log.Panicf("error creating router: %s", err)
	}

	return http.ListenAndServe(fmt.Sprintf("%s:%s", srv.Config.AppHost, srv.Config.AppPort), router)
}

func (srv *Server) createRouter() (http.Handler, error) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	handler, err := srv.createHandler()
	if err != nil {
		return nil, err
	}

	setupRoutes(router, handler)
	return router, nil
}

func (srv *Server) createHandler() (*apiv1.Handler, error) {
	db, err := gorm.Open(mysql.Open(srv.Config.DBConnString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	c := cache.New(20*time.Minute, 10*time.Minute)

	handler := &apiv1.Handler{
		Environment: srv.Config.Environment,
		DB:          db,
		C:           c,
	}

	return handler, nil
}

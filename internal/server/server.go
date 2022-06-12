package server

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/C22-PS350/backend-rawati/internal/server/apiv1"
	"github.com/C22-PS350/backend-rawati/internal/utils"
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
		return err
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

	rand.Seed(time.Now().UnixNano())
	handler, err := srv.createHandler()
	if err != nil {
		return nil, err
	}

	setupRoutes(router, handler)
	return router, nil
}

func (srv *Server) createHandler() (*apiv1.Handler, error) {
	dbPool, err := sql.Open("mysql", srv.Config.DBConnString)
	if err != nil {
		return nil, err
	}

	if err := dbPool.Ping(); err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbPool,
	},
	), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	c := cache.New(20*time.Minute, 10*time.Minute)
	v := NewValidator()

	r := &apiv1.Refs{
		Environment: srv.Config.Environment,
		ModelAPIUrl: srv.Config.ModelAPIUrl,
	}

	d := &apiv1.Deps{
		DB: db,
		C:  c,
		V:  v,
	}

	handler := &apiv1.Handler{
		Refs: r,
		Deps: d,
	}

	if srv.Config.Environment == utils.Remote {
		handler, err = srv.setupGCPClients(handler)
		if err != nil {
			return nil, err
		}
	}

	return handler, nil
}

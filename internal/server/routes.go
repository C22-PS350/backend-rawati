package server

import (
	"github.com/farryl/project-mars/internal/server/apiv1"
	"github.com/go-chi/chi/v5"
)

func setupRoutes(r *chi.Mux, h *apiv1.Handler) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/hello", h.Hello)
		r.Post("/users", h.CreateUser)

		// r.Group(func(r chi.Router) {
		// 	r.Use(h.AuthMiddleware)
		// })
	})
}

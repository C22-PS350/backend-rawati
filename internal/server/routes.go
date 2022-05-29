package server

import (
	"github.com/C22-PS350/backend-rawati/internal/server/apiv1"
	"github.com/go-chi/chi/v5"
)

func setupRoutes(r *chi.Mux, h *apiv1.Handler) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", h.Register)
		r.Post("/auth/login", h.Login)

		// r.Group(func(r chi.Router) {
		// 	r.Use(h.AuthMiddleware)
		// })
	})
}

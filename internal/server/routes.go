package server

import (
	"github.com/C22-PS350/backend-rawati/internal/server/apiv1"
	"github.com/go-chi/chi/v5"
)

func setupRoutes(r *chi.Mux, h *apiv1.Handler) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", h.Register)
		r.Post("/auth/login", h.Login)
		r.Put("/auth/forgot-password", h.ForgotPassword)

		r.Group(func(r chi.Router) {
			r.Use(h.AuthMiddleware)
			r.Get("/users/{user_id}", h.GetUser)
			r.Put("/users/{user_id}/update", h.UpdateUser)
			r.Put("/users/{user_id}/update-password", h.UpdateUserPassword)
			r.Delete("/users/{user_id}", h.DeleteUser)
		})
	})
}

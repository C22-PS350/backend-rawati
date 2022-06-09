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
			// user
			r.Get("/users/{user_id}", h.GetUser)
			r.Put("/users/{user_id}", h.UpdateUser)
			r.Put("/users/{user_id}/update-password", h.UpdateUserPassword)
			r.Delete("/users/{user_id}", h.DeleteUser)

			// user profile
			r.Get("/users/{user_id}/profile", h.GetUserProfile)
			r.Put("/users/{user_id}/profile", h.UpdateUserProfile)
			// r.Post("/users/{user_id}/profile", h.CreateUserProfile)

			// activiy (exercise)
			r.Get("/users/{user_id}/exercises", h.GetAllExerciseActivity)
			r.Get("/users/{user_id}/exercises/{exercise_id}", h.GetExerciseActivity)
			// r.Post("/users/{user_id}/exercises", h.CreateExerciseActivity)

			// activiy (food)
			r.Get("/users/{user_id}/foods", h.GetAllFoodActivity)
			r.Get("/users/{user_id}/foods/{food_id}", h.GetFoodActivity)
			r.Post("/users/{user_id}/foods", h.CreateFoodActivity)

			// recommendation (exercise)
			r.Post("/recommendation/exercise", h.CreateExerciseRecommendation)

			// recommendation (food)
			r.Post("/recommendation/food", h.CreateFoodRecommendation)

			// resources
			r.Get("/exercises", h.GetAllExercises)
			r.Get("/foods", h.GetAllFoods)
		})
	})
}

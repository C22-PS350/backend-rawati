package models

type ExerciseRecommendationRequest struct {
	Calories float64 `json:"calories" validate:"required,number" example:"200"`
}

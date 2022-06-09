package models

type FoodRecommendationRequest struct {
	Calories float64 `json:"calories" validate:"required,number"`
}

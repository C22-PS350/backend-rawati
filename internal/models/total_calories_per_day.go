package models

type TotalCaloriesPerDayResponse struct {
	UserID        uint64  `json:"user_id"`
	ExerciseTotal float64 `json:"exercise_total"`
	FoodTotal     float64 `json:"food_total"`
}

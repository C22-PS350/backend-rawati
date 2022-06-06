package models

type ExerciseResponse struct {
	ExerciseID uint64 `json:"exercise_id"`
	Name       string `json:"name"`
	Calories   uint64 `json:"calories"`
}

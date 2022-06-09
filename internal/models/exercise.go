package models

type ExerciseResponse struct {
	ExerciseID uint64  `json:"exercise_id"`
	Name       string  `json:"name"`
	Calories   float64 `json:"calories"`
}

type ExerciseTest1 struct {
	ExerciseID uint64  `faker:"unique"`
	Name       string  `faker:"-"`
	Calories   float64 `faker:"-"`
}

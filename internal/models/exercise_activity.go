package models

import "time"

type ExerciseActivityRequest struct {
	ExerciseActivityID uint64    `json:"-" gorm:"primaryKey"`
	UserID             uint64    `json:"-"`
	Name               string    `json:"name" validate:"required" maxLength:"60" example:"voleey"`
	ExerciseDate       time.Time `json:"-"`
	Duration           int       `json:"duration" example:"60"`
	Calories           float64   `json:"calories" validate:"required" example:"200"`
}

type ExerciseActivityCreateResponse struct {
	ExerciseActivityID uint64 `json:"exercise_activity_id"`
	UserID             uint64 `json:"user_id"`
}

type ExerciseActivityResponse struct {
	ExerciseActivityID uint64     `json:"exercise_activity_id"`
	UserID             uint64     `json:"user_id"`
	Name               string     `json:"name"`
	ExerciseDate       *time.Time `json:"exercise_date"`
	Duration           int        `json:"duration"`
	Calories           float64    `json:"calories"`
}

type ExerciseActivityTest1 struct {
	UserID     uint64 `faker:"unique"`
	Name       string `faker:"name"`
	Username   string `faker:"username,unique"`
	Email      string `faker:"email,unique"`
	Password   string `faker:"len=30"`
	IsVerified bool
}

type ExerciseActivityTest2 struct {
	ExerciseActivityID uint64    `faker:"unique"`
	UserID             uint64    `faker:"-"`
	Name               string    `faker:"name"`
	ExerciseDate       time.Time `faker:"-"`
	Duration           int       `faker:"boundary_start=0, boundary_end=120"`
	Calories           float64   `faker:"boundary_start=0, boundary_end=500"`
}

package models

import "time"

type ExerciseActivityRequest struct {
	ExerciseActivityID uint64    `json:"-" gorm:"primaryKey"`
	UserID             uint64    `json:"-"`
	Name               string    `json:"name" validate:"required,max=60" maxLength:"60" example:"voleey"`
	ExerciseDate       time.Time `json:"-"`
	Duration           float64   `json:"duration" validate:"required,number" example:"60"`
	HeartRate          float64   `json:"heart_rate" gorm:"-" validate:"required,number" example:"100"`
	BodyTemp           float64   `json:"body_temp" gorm:"-" example:"36.7"`
	Calories           float64   `json:"-" example:"200"`
}

type ExerciseActivityUserData struct {
	Height    float64   `validate:"required"`
	Weight    float64   `validate:"required"`
	Gender    string    `validate:"required"`
	BirthDate time.Time `validate:"required"`
}

type ExerciseActivityPredictRequest struct {
	Gender    int     `json:"gender"`
	Age       int     `json:"age"`
	Weight    float64 `json:"weight"`
	Height    float64 `json:"height"`
	Duration  float64 `json:"duration"`
	HeartRate float64 `json:"heart_rate"`
	BodyTemp  float64 `json:"body_temp,omitempty"`
}

type ExerciseActivityPredictResult struct {
	Calories float64 `json:"calories"`
}

type ExerciseActivityCreateResponse struct {
	ExerciseActivityID uint64  `json:"exercise_activity_id"`
	UserID             uint64  `json:"user_id"`
	Calories           float64 `json:"calories"`
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

type ExerciseActivityTest3 struct {
	ExerciseActivityID uint64    `faker:"-"`
	UserID             uint64    `faker:"-"`
	Name               string    `faker:"name"`
	ExerciseDate       time.Time `faker:"-"`
	Duration           float64   `faker:"boundary_start=30, boundary_end=90"`
	HeartRate          float64   `faker:"boundary_start=70, boundary_end=110"`
	BodyTemp           float64   `faker:"boundary_start=35, boundary_end=38"`
	Calories           float64   `faker:"-"`
}

type ExerciseActivityTest4 struct {
	ProfileID  uint64 `faker:"unique"`
	UserID     uint64 `faker:"-"`
	Height     uint16 `faker:"boundary_start=0,boundary_end=200"`
	Weight     uint16 `faker:"boundary_start=0,boundary_end=200"`
	WeightGoal uint16 `faker:"boundary_start=0,boundary_end=80"`
	Gender     string `faker:"oneof: L, P"`
	BirthDate  time.Time
}

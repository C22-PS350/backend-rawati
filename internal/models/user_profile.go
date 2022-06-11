package models

import "time"

type GetUserProfileResponse struct {
	ProfileID  uint64     `json:"profile_id"`
	UserID     uint64     `json:"user_id"`
	Height     uint16     `json:"height"`
	Weight     uint16     `json:"weight"`
	WeightGoal uint16     `json:"weight_goal"`
	Gender     string     `json:"gender"`
	BirthDate  *time.Time `json:"birth_date"`
}

// create or update
type UserProfileRequest struct {
	ProfileID  uint64     `json:"-" gorm:"primaryKey"`
	UserID     *uint64    `json:"-"`
	Height     *uint16    `json:"height" example:"183"`
	Weight     *uint16    `json:"weight" example:"70"`
	WeightGoal *uint16    `json:"weight_goal" example:"60"`
	Gender     *string    `json:"gender" example:"L (male) or P (female)"`
	BirthDate  *time.Time `json:"birth_date" example:"2022-06-10T14:39:11Z (ISO 8601 with UTC timezone)"`
}

type UserProfileResponse struct {
	ProfileID uint64 `json:"profile_id"`
	UserID    uint64 `json:"user_id"`
}

// test
type UserProfileTest1 struct {
	UserID     uint64 `faker:"unique"`
	Name       string `faker:"name"`
	Username   string `faker:"username,unique"`
	Email      string `faker:"email,unique"`
	Password   string `faker:"len=30"`
	IsVerified bool
}

type UserProfileTest2 struct {
	ProfileID  uint64 `faker:"unique"`
	UserID     uint64 `faker:"-"`
	Height     uint16 `faker:"boundary_start=0,boundary_end=200"`
	Weight     uint16 `faker:"boundary_start=0,boundary_end=200"`
	WeightGoal uint16 `faker:"boundary_start=0,boundary_end=80"`
	Gender     string `faker:"oneof: L, P"`
	BirthDate  time.Time
}

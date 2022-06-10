package models

import "time"

type FoodActivityRequest struct {
	FoodActivityID uint64    `json:"-" gorm:"primaryKey"`
	UserID         uint64    `json:"-"`
	Name           string    `json:"name" validate:"required" maxLength:"60" example:"leiyker"`
	FoodDate       time.Time `json:"-"`
	Calories       float64   `json:"calories" validate:"required" example:"200"`
}

type FoodActivityCreateResponse struct {
	FoodActivityID uint64 `json:"food_activity_id"`
	UserID         uint64 `json:"user_id"`
}

type FoodActivityResponse struct {
	FoodActivityID uint64     `json:"food_activity_id"`
	UserID         uint64     `json:"user_id"`
	Name           string     `json:"name"`
	FoodDate       *time.Time `json:"food_date"`
	Calories       float64    `json:"calories"`
}

type FoodActivityTest1 struct {
	UserID     uint64 `faker:"unique"`
	Name       string `faker:"name"`
	Username   string `faker:"username,unique"`
	Email      string `faker:"email,unique"`
	Password   string `faker:"len=30"`
	IsVerified bool
}

type FoodActivityTest2 struct {
	FoodActivityID uint64    `faker:"unique"`
	UserID         uint64    `faker:"-"`
	Name           string    `faker:"name"`
	FoodDate       time.Time `faker:"-"`
	Calories       float64   `faker:"boundary_start=0, boundary_end=500"`
}

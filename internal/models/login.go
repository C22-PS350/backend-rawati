package models

import "time"

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required" example:"farrylvanhouten (username) or farryl@gmail.com (email)"`
	Password   string `json:"password" validate:"required" example:"kmzwa8awaa"`
}

type LoginData struct {
	UserID   uint64 `validate:"required"`
	Password string `validate:"required"`
	Token    string `validate:"required"`
}

type LoginResponse struct {
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

type LoginTest1 struct {
	UserID     uint64 `faker:"unique"`
	Name       string `faker:"name"`
	Username   string `faker:"username,unique"`
	Email      string `faker:"email,unique"`
	Password   string `faker:"len=30"`
	IsVerified bool
}

type LoginTest2 struct {
	UserID    uint64    `faker:"-"`
	Token     string    `faker:"len=40"`
	CreatedAt time.Time `faker:"-"`
}

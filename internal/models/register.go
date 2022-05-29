package models

import "time"

type RegisterRequest struct {
	UserID   uint32  `json:"-" gorm:"primaryKey" swaggerignore:"true"`
	Name     *string `json:"name" validate:"required,max=60"`
	Username *string `json:"username" validate:"required,alphanum,max=30"`
	Email    *string `json:"email" validate:"required,email,max=60"`
	Password *string `json:"password" validate:"required,max=60"`
}

type RegisterUserToken struct {
	UserID    uint32
	Token     string
	CreatedAt time.Time
}

type RegisterResponse struct {
	UserID uint32 `json:"user_id"`
}

type RegisterTest1 struct {
	Name     string `faker:"name"`
	Username string `faker:"username,unique"`
	Email    string `faker:"email,unique"`
	Password string `faker:"len=30"`
}

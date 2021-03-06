package models

import "time"

type RegisterRequest struct {
	UserID     uint64  `json:"-" gorm:"primaryKey"`
	Name       *string `json:"name" validate:"required,max=60" maxLength:"60" example:"farryl van houten"`
	Username   *string `json:"username" validate:"required,alphanum,max=30" maxLength:"30" example:"farrylvanhouten"`
	Email      *string `json:"email" validate:"required,email,max=60" maxLength:"60" example:"farryl@gmail.com"`
	Password   *string `json:"password" validate:"required,max=60" maxLength:"60" example:"kmzwa8awaa"`
	IsVerified bool    `json:"-"`
}

type RegisterUserToken struct {
	UserID    uint64
	Token     string
	CreatedAt time.Time
}

type RegisterResponse struct {
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterTest1 struct {
	Name     string `faker:"name"`
	Username string `faker:"username,unique"`
	Email    string `faker:"email,unique"`
	Password string `faker:"len=30"`
}

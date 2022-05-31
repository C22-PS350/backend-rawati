package models

// GetUser
type GetUserResponse struct {
	UserID     uint64 `json:"user_id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

// UpdateUser
type UpdateUserRequest struct {
	UserID   uint64  `json:"-" gorm:"primaryKey"`
	Name     *string `json:"name" validate:"required,max=60"`
	Username *string `json:"username" validate:"required,alphanum,max=30"`
	Email    *string `json:"email" validate:"required,email,max=60"`
}

type UpdateUserResponse struct {
	UserID uint64 `json:"user_id"`
}

type UpdateUserTest1 struct {
	UserID     uint64 `faker:"unique"`
	Name       string `faker:"name"`
	Username   string `faker:"username,unique"`
	Email      string `faker:"email,unique"`
	Password   string `faker:"len=30"`
	IsVerified bool
}

// UpdateUserPassword
type UpdateUserPwdRequest struct {
	OldPassword string `json:"old_password" validate:"required,max=60"`
	NewPassword string `json:"new_password" validate:"required,max=60"`
}

type UpdateUserPwdData struct {
	Password string `validate:"required"`
}

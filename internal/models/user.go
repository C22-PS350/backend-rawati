package models

type User struct {
	UserID   *uint32 `json:"user_id" swaggerignore:"true"`
	Name     *string `json:"name"`
	Password *string `json:"password"`
}

type UserResponse struct {
	UserID   uint32 `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

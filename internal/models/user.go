package models

type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserID   uint32 `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

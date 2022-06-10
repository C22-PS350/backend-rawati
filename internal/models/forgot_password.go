package models

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email,max=60" example:"farryl@gmail.com"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

type ForgotPasswordData struct {
	UserID   uint64
	Username string
	Email    string
}

type ForgotPasswordMessage struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPasswordTest1 struct {
	UserID     uint64 `faker:"unique"`
	Name       string `faker:"name"`
	Username   string `faker:"username,unique"`
	Email      string `faker:"email,unique"`
	Password   string `faker:"len=30"`
	IsVerified bool
}

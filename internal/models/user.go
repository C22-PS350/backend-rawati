package models

type User struct {
	UserID   uint32 `gorm:"primaryKey"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

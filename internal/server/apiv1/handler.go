package apiv1

import (
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}
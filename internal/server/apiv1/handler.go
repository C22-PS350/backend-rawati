package apiv1

import (
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
	C  *cache.Cache
}

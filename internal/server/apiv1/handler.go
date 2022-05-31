package apiv1

import (
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type Handler struct {
	Environment string
	DB          *gorm.DB
	C           *cache.Cache
}

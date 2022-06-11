package apiv1

import (
	"cloud.google.com/go/pubsub"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type Handler struct {
	Environment string
	ModelAPIUrl string
	DB          *gorm.DB
	C           *cache.Cache
	GcpClient   *GcpClient
}

type GcpClient struct {
	PubSub *pubsub.Client
}

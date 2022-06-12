package apiv1

import (
	"cloud.google.com/go/pubsub"
	"github.com/go-playground/validator/v10"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type Handler struct {
	*Refs
	*Deps
	*GcpClients
}

type Refs struct {
	Environment string
	ModelAPIUrl string
}

type Deps struct {
	DB *gorm.DB
	C  *cache.Cache
	V  *validator.Validate
}

type GcpClients struct {
	PubSub *pubsub.Client
}

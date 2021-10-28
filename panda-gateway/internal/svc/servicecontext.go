package svc

import (
	"Pandax/panda-gateway/internal/config"
	"Pandax/panda-gateway/internal/middleware"
	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	CheckLogin rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CheckLogin: middleware.NewCheckLoginMiddleware().Handle,
	}
}

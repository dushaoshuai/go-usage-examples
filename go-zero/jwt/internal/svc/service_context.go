package svc

import (
	"github.com/dushaoshuai/go-usage-examples/go-zero/jwt/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

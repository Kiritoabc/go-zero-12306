package svc

import (
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

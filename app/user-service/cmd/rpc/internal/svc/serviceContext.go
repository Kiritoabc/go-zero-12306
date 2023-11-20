package svc

import (
	"go-zero-12306/app/user-service/cmd/rpc/internal/config"
	"go-zero-12306/app/user-service/model/tUser"
)

type ServiceContext struct {
	Config     config.Config
	User0Model tUser.TUser0Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

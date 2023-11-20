package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-12306/app/user-service/cmd/api/internal/config"
	"go-zero-12306/app/user-service/cmd/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}

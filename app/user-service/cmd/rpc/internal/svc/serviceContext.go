package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/user-service/cmd/rpc/internal/config"
	"go-zero-12306/app/user-service/model/tUser"
)

type ServiceContext struct {
	Config     config.Config
	User0Model tUser.TUser0Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlCon := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:     c,
		User0Model: tUser.NewTUser0Model(sqlCon, c.Cache),
	}
}

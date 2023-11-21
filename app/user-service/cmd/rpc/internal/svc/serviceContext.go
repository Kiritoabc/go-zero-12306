package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/user-service/cmd/rpc/internal/config"
	"go-zero-12306/app/user-service/model/tUser"
)

type ServiceContext struct {
	Config      config.Config
	User0Model  tUser.TUser0Model
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlCon := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass, // no password set
			DB:       0,            // use default DB
		}),
		User0Model: tUser.NewTUser0Model(sqlCon, c.Cache),
	}
}

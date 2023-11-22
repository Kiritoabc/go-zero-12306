package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/user-service/cmd/rpc/internal/config"
	"go-zero-12306/app/user-service/model/tPassenger"
	"go-zero-12306/app/user-service/model/tUser"
	"go-zero-12306/app/user-service/model/tUserDeletion"
	"go-zero-12306/app/user-service/model/tUserMail"
	"go-zero-12306/app/user-service/model/tUserPhone"
	"go-zero-12306/app/user-service/model/tUserReuse"
)

type ServiceContext struct {
	Config            config.Config
	RedisClient       *redis.Client
	Passenger0Model   tPassenger.TPassenger0Model
	User0Model        tUser.TUser0Model
	UserDeletionModel tUserDeletion.TUserDeletionModel
	UserMail0Model    tUserMail.TUserMail0Model
	UserPhone0Model   tUserPhone.TUserPhone0Model
	UserReuseModel    tUserReuse.TUserReuseModel
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
		Passenger0Model:   tPassenger.NewTPassenger0Model(sqlCon, c.Cache),
		User0Model:        tUser.NewTUser0Model(sqlCon, c.Cache),
		UserDeletionModel: tUserDeletion.NewTUserDeletionModel(sqlCon, c.Cache),
		UserMail0Model:    tUserMail.NewTUserMail0Model(sqlCon, c.Cache),
		UserPhone0Model:   tUserPhone.NewTUserPhone0Model(sqlCon, c.Cache),
		UserReuseModel:    tUserReuse.NewTUserReuseModel(sqlCon, c.Cache),
	}
}

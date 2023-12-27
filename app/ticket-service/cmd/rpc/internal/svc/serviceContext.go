package svc

import (
	"github.com/go-redis/redis/v8"
	go_redis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/config"
	"go-zero-12306/app/ticket-service/model/tCarriage"
	"go-zero-12306/app/ticket-service/model/tRegion"
	"go-zero-12306/app/ticket-service/model/tSeat"
	"go-zero-12306/app/ticket-service/model/tStation"
	"go-zero-12306/app/ticket-service/model/tTicket"
	"go-zero-12306/app/ticket-service/model/tTrain"
	"go-zero-12306/app/ticket-service/model/tTrainStation"
	"go-zero-12306/app/ticket-service/model/tTrainStationPrice"
	"go-zero-12306/app/ticket-service/model/tTrainStationRelation"
)

type ServiceContext struct {
	Config                     config.Config
	RedisClient                *redis.Client
	RedisClient1               *go_redis.Redis
	TCarriageModel             tCarriage.TCarriageModel
	TRegionModel               tRegion.TRegionModel
	TSeatModel                 tSeat.TSeatModel
	TStationModel              tStation.TStationModel
	TTicketModel               tTicket.TTicketModel
	TTrainModel                tTrain.TTrainModel
	TTrainStationModel         tTrainStation.TTrainStationModel
	TTrainStationPriceModel    tTrainStationPrice.TTrainStationPriceModel
	TTrainStationRelationModel tTrainStationRelation.TTrainStationRelationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlCon := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass, // no password set
			DB:       1,            // use default DB
		}),
		RedisClient1: go_redis.New(c.Redis.Host, func(r *go_redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		TCarriageModel:             tCarriage.NewTCarriageModel(sqlCon, c.Cache),
		TRegionModel:               tRegion.NewTRegionModel(sqlCon, c.Cache),
		TSeatModel:                 tSeat.NewTSeatModel(sqlCon, c.Cache),
		TStationModel:              tStation.NewTStationModel(sqlCon, c.Cache),
		TTicketModel:               tTicket.NewTTicketModel(sqlCon, c.Cache),
		TTrainModel:                tTrain.NewTTrainModel(sqlCon, c.Cache),
		TTrainStationModel:         tTrainStation.NewTTrainStationModel(sqlCon, c.Cache),
		TTrainStationPriceModel:    tTrainStationPrice.NewTTrainStationPriceModel(sqlCon, c.Cache),
		TTrainStationRelationModel: tTrainStationRelation.NewTTrainStationRelationModel(sqlCon, c.Cache),
	}
}

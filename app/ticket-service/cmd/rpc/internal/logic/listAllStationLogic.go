package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-12306/common/constant"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAllStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllStationLogic {
	return &ListAllStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListAllStationLogic) ListAllStation(in *pb.ListAllStationReq) (*pb.ListAllStationResp, error) {
	redisValue := l.svcCtx.RedisClient.Get(l.ctx, constant.STATION_ALL)
	value := redisValue.Val()
	var resp = &pb.ListAllStationResp{}
	// 缓存存在
	if value != "" {
		err := json.Unmarshal([]byte(value), &resp.StationQueryRespDTO)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}

	// 上锁
	lock := redis.NewRedisLock(l.svcCtx.RedisClient1, constant.SAFE_GET_DISTRIBUTED_LOCK_KEY_PREFIX+constant.STATION_ALL)
	// 设置过期时间
	lock.SetExpire(10)
	// 尝试获取锁
	acquire, err := lock.Acquire()

	switch {
	case err != nil:
		return resp, err
	case acquire:
		// 获取到锁
		defer lock.Release() // 释放锁
		// 缓存不存在查询缓存
		builder := l.svcCtx.TStationModel.SelectBuilder()
		list, err := l.svcCtx.TStationModel.List(l.ctx, builder, "")
		if err != nil {
			return &pb.ListAllStationResp{}, err
		}
		copier.Copy(&resp.StationQueryRespDTO, &list)
		// 存入缓存
		cacheValue, err := json.Marshal(&resp.StationQueryRespDTO)
		if err != nil {
			return resp, err
		}
		//err = l.svcCtx.RedisClient1.Set(constant.STATION_ALL, string(cacheValue))
		// 15 天
		//err = l.svcCtx.RedisClient1.Setex(constant.STATION_ALL, string(cacheValue), 15*24*60*60)
		err = l.svcCtx.RedisClient.Set(l.ctx, constant.STATION_ALL, string(cacheValue), constant.ADVANCE_TICKET_DAY).Err()
		if err != nil {
			return resp, err
		}
		return resp, nil
	case !acquire:

	}
	return &pb.ListAllStationResp{}, nil
}

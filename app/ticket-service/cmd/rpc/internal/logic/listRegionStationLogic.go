package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/regionStationQueryTypeEnum"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"
	"go-zero-12306/common/constant"
	"strconv"
)

type ListRegionStationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRegionStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRegionStationLogic {
	return &ListRegionStationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * 查询车站&城市站点集合信息
 *
 * @param  车站&站点查询参数
 * @return 车站&站点返回数据集合
 */

func (l *ListRegionStationLogic) ListRegionStation(in *pb.ListRegionStationReq) (*pb.ListRegionStationResp, error) {

	var key string
	if in.Name != "" {
		key = constant.REGION_STATION + in.Name
		// 查询
	}

	key = constant.REGION_STATION + strconv.FormatInt(in.QueryType, 10)
	// params
	// 热门标识
	var popularFlag int64
	// 地区首字母
	var initialList []string
	switch in.QueryType {
	case 0:
		popularFlag = 1
	case 1:
		initialList = regionStationQueryTypeEnum.RegionstationquerytypeenumA_Espells
	case 2:
		initialList = regionStationQueryTypeEnum.RegionstationquerytypeenumF_Jspells
	case 3:
		initialList = regionStationQueryTypeEnum.RegionstationquerytypeenumK_Ospells
	case 4:
		initialList = regionStationQueryTypeEnum.RegionstationquerytypeenumP_Tspells
	case 5:
		initialList = regionStationQueryTypeEnum.RegionstationquerytypeenumU_Zspells
	default:
		return &pb.ListRegionStationResp{}, errors.Wrapf(errors.New("查询失败，请检查查询参数是否正确"), "查询失败，请检查查询参数是否正确")
	}

	var resp = &pb.ListRegionStationResp{}
	// 查询缓存
	redisValue := l.svcCtx.RedisClient.Get(l.ctx, key)
	value := redisValue.Val()

	// 存在，返回
	if value != "" {
		err := json.Unmarshal([]byte(value), &resp.TrainStationQueryRespDTO)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
	// 上锁
	lock := redis.NewRedisLock(l.svcCtx.RedisClient1, fmt.Sprintf(constant.LOCK_QUERY_REGION_STATION_LIST, in.QueryType))
	// 设置过期时间
	lock.SetExpire(10)
	// 尝试获取锁
	acquire, err := lock.Acquire()
	switch {
	case err != nil:
		return resp, err
	case acquire:
		defer lock.Release()
		// todo: 获取到锁，查询数据库(抽离出来)
		_ = popularFlag
		_ = initialList
	case !acquire:
		// await
	}

	return resp, nil
}

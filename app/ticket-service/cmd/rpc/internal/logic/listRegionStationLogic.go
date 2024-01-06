package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
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
	// name 存在，根据name查询
	if in.Name != "" {
		key = constant.REGION_STATION + in.Name
		builder := l.svcCtx.TStationModel.SelectBuilder()

		return l.safeGetRegionStation(key, func() (*pb.ListRegionStationResp, error) {
			var resp = &pb.ListRegionStationResp{}
			tStations, err := l.svcCtx.TStationModel.SelectListByName(l.ctx, builder, in.Name)
			if err != nil {
				return nil, err
			}
			err = copier.Copy(&resp.RegionStationQueryRespDTO, &tStations)
			if err != nil {
				return resp, err
			}
			return resp, nil
		}, in.Name)
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

	return l.safeGetRegionStation(key, func() (*pb.ListRegionStationResp, error) {
		var resp = &pb.ListRegionStationResp{}
		builder := l.svcCtx.TRegionModel.SelectBuilder()
		tRegions, err := l.svcCtx.TRegionModel.SelectListByInitialOrPopularFlag(l.ctx, builder, popularFlag, initialList)
		if err != nil {
			return resp, err
		}
		err = copier.Copy(&resp.RegionStationQueryRespDTO, &tRegions)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}, strconv.FormatInt(in.QueryType, 10))
}

func (l ListRegionStationLogic) safeGetRegionStation(key string, fn func() (*pb.ListRegionStationResp, error), params string) (*pb.ListRegionStationResp, error) {
	var resp = &pb.ListRegionStationResp{}
	// 1.查缓存
	redisValue := l.svcCtx.RedisClient.Get(l.ctx, key).Val()
	if redisValue != "" {
		err := copier.Copy(&resp.RegionStationQueryRespDTO, &redisValue)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
	// 上锁
	lock := redis.NewRedisLock(l.svcCtx.RedisClient1, fmt.Sprintf(constant.LOCK_QUERY_REGION_STATION_LIST, params))
	// 设置过期时间
	lock.SetExpire(10)
	// 尝试获取锁
	acquire, err := lock.Acquire()
	switch {
	case err != nil:
		return resp, err
	case acquire:
		defer lock.Release()
		// 查询数据
		fnResp, err := fn()
		if err != nil {
			return resp, err
		}
		resp.RegionStationQueryRespDTO = fnResp.RegionStationQueryRespDTO
		//存入缓存
		cacheValue, err := json.Marshal(&resp.RegionStationQueryRespDTO)
		err = l.svcCtx.RedisClient.Set(l.ctx, key, cacheValue, constant.ADVANCE_TICKET_DAY).Err()
		if err != nil {
			return resp, err
		}
		return resp, nil
	case !acquire:
		// await
	}
	return resp, nil
}

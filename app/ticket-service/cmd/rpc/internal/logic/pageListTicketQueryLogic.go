package logic

import (
	"context"
	"encoding/json"
	"fmt"
	godisson "github.com/cheerego/go-redisson"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/dto"
	"go-zero-12306/common/constant"
	"sort"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageListTicketQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageListTicketQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageListTicketQueryLogic {
	return &PageListTicketQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TODO: TicketControllerRpc
func (l *PageListTicketQueryLogic) PageListTicketQuery(in *pb.PageListTicketQueryReq) (*pb.PageListTicketQueryResp, error) {
	// todo: 责任链模式？？（纯nt）参数不考虑
	var resp *pb.PageListTicketQueryResp
	// 1. 从redis中查询到数据
	redisClient := l.svcCtx.RedisClient
	hmGet := redisClient.HMGet(l.ctx, constant.REGION_TRAIN_STATION_MAPPING, []string{in.FromStation, in.ToStation}...)
	if hmGet.Err() != nil {
		return resp, hmGet.Err()
	}
	stationDetails := hmGet.Val()
	count := NotNilCount(stationDetails)

	// 2. count > 0?
	if count > 0 {
		// redisson 上锁
		lock := godisson.NewGodisson(redisClient).NewRLock(constant.LOCK_REGION_TRAIN_STATION_MAPPING)
		lock.Lock()
		// 再次查询
		stationDetails = redisClient.
			HMGet(l.ctx, constant.REGION_TRAIN_STATION_MAPPING, []string{in.FromStation, in.ToStation}...).
			Val()
		count = NotNilCount(stationDetails)
		if count > 0 {
			// 查询所有 station
			builder := l.svcCtx.TStationModel.SelectBuilder()
			stationDOList, _ := l.svcCtx.TStationModel.List(l.ctx, builder, "")
			var regionTrainStationMap = map[string]string{}
			for _, stationDO := range stationDOList {
				regionTrainStationMap[stationDO.Code.String] = stationDO.RegionName.String
			}
			// 将map转换为[key, value]对切片，以便于HMSet方法使用
			//fieldValuePairs := make([]interface{}, 0, len(regionTrainStationMap)*2)
			//for k, v := range regionTrainStationMap {
			//	fieldValuePairs = append(fieldValuePairs, k, v)
			//}
			//redisClient.HMSet(l.ctx, constant.REGION_TRAIN_STATION_MAPPING, fieldValuePairs...)
			redisClient.HMSet(l.ctx, constant.REGION_TRAIN_STATION_MAPPING, regionTrainStationMap)
			stationDetails = make([]interface{}, 0, 2)
			stationDetails = append(stationDetails, in.FromStation, in.ToStation)

			// 释放锁
			lock.Unlock()
		}
	}
	// 3.查询座位
	seatResults := make([]*dto.TicketListDTO, 0)
	buildRegionTrainStationHashKey := fmt.Sprintf(constant.REGION_TRAIN_STATION,
		stationDetails[0], stationDetails[1])
	hashValues, _ := redisClient.HGetAll(l.ctx, buildRegionTrainStationHashKey).Result()
	// 将redis-go返回的map[string]string直接转换为所需类型
	regionTrainStationAllMap := make(map[string]string)
	for k, v := range hashValues {
		regionTrainStationAllMap[k] = v
	}
	if regionTrainStationAllMap == nil || len(regionTrainStationAllMap) == 0 {
		// 上锁
		lock := godisson.NewGodisson(redisClient).NewRLock(constant.LOCK_REGION_TRAIN_STATION)
		lock.Lock()
		// 再次查询
		regionTrainStationAllMap, _ = redisClient.HGetAll(l.ctx, buildRegionTrainStationHashKey).Result()
		if regionTrainStationAllMap == nil || len(regionTrainStationAllMap) == 0 {
			// 查询数据库

		}
		lock.Unlock()
	}

	// 4. 给seat赋值
	if len(seatResults) == 0 {
		// 赋值给seatResults
		// 将regionTrainStationAllMap的值（JSON字符串）反序列化为TicketListDTO对象
		var results []string
		for _, each := range regionTrainStationAllMap {
			results = append(results, each)
		}

		for _, jsonStr := range results {
			result := new(dto.TicketListDTO)
			json.Unmarshal([]byte(jsonStr), &result)
			seatResults = append(seatResults, result)
		}
	}
	// 确保seatResults已排序
	sort.Slice(seatResults, func(i, j int) bool {
		// 这里根据需要实现排序逻辑，例如按时间排序
		return seatResults[i].DepartureTime < seatResults[j].DepartureTime
	})

	// 5. 存入缓存

	return &pb.PageListTicketQueryResp{}, nil
}

func NotNilCount(in []interface{}) int64 {
	var count int64 = 0
	for _, value := range in {
		if value != nil {
			count++
		}
	}

	return count
}

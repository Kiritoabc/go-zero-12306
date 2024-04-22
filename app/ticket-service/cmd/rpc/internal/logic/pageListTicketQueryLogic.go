package logic

import (
	"context"
	"encoding/json"
	"fmt"
	godisson "github.com/cheerego/go-redisson"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/dto"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"
	"go-zero-12306/app/ticket-service/model/tTrain"
	"go-zero-12306/common/constant"
	"go-zero-12306/common/tool"
	"go-zero-12306/common/tool/cacheUtil"
	"sort"
	"strconv"
	"strings"

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
		err := lock.Lock()
		if err != nil {
			return resp, err
		}
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

	// 获取 distributedCache
	distributedCache := NewDistributedCacheLogic(l.ctx, l.svcCtx)

	// 3.查询列车座位信息
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
			builder := l.svcCtx.TTrainStationRelationModel.SelectBuilder()
			trainStationRelationList, err := l.svcCtx.TTrainStationRelationModel.SelectList(l.ctx, builder, "")
			if err != nil {
				return resp, err
			}
			for _, trainStationRelation := range trainStationRelationList {
				// 将json字符串反序列化为TicketListDTO对象
				fmt.Sprintf("%v", trainStationRelation)
				var trainDO tTrain.TTrain
				// safeGet() 获取缓存
				key := constant.TRAIN_INFO + strconv.FormatInt(trainStationRelation.TrainId.Int64, 10)

				target, err := distributedCache.SafeGet(key,
					trainDO,
					nil,
					func() (interface{}, error) {
						return nil, nil
					}, nil,
					constant.ADVANCE_TICKET_DAY)
				if err != nil {
					return resp, err
				}
				trainDO = target.(tTrain.TTrain)

				result := &dto.TicketListDTO{}
				result.TrainId = strconv.Itoa(int(trainDO.Id))
				result.TrainNumber = trainDO.TrainNumber.String
				result.DepartureTime = tool.ConvertDateToLocalTime(trainStationRelation.DepartureTime.Time.String(), "HH:mm")
				result.ArrivalTime = tool.ConvertDateToLocalTime(trainStationRelation.ArrivalTime.Time.String(), "HH:mm")
				result.Duration = strconv.Itoa(tool.CalculateHourDifference(trainStationRelation.DepartureTime.Time.String(), trainStationRelation.ArrivalTime.Time.String()))
				result.Departure = trainStationRelation.Departure.String
				result.Arrival = trainStationRelation.Arrival.String
				result.DepartureFlag = trainStationRelation.DepartureFlag.Int64 == 1
				result.ArrivalFlag = trainStationRelation.ArrivalFlag.Int64 == 1
				result.TrainType = int(trainDO.TrainType.Int64)
				result.TrainBrand = trainDO.TrainBrand.String

				if !trainDO.TrainTag.Valid {
					result.TrainTags = strings.Split(trainDO.TrainTag.String, ",")
				}
				//betweenDay := dateutil.DateDiff("day", trainDO.DepartureTime, trainDO.ArrivalTime, false)

				buildKey, _ := cacheUtil.BuildKey([]string{
					strconv.FormatInt(trainStationRelation.TrainId.Int64, 10),
					trainStationRelation.Departure.String,
					trainStationRelation.Departure.String})
				resultBytes, _ := json.Marshal(result)
				regionTrainStationAllMap[buildKey] = string(resultBytes)
			}
			// putAll() 存入缓存
			l.svcCtx.RedisClient.HMSet(l.ctx, buildRegionTrainStationHashKey, regionTrainStationAllMap)
		}
		lock.Unlock()
	}

	// 4. 给列车seat信息赋值票价和多少张
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
	for _, result := range seatResults {
		// safeGet() 从缓存中获取price数据

		// load price to seat

		// cache get 站点余票查询 TRAIN_STATION_REMAINING_TICKET, 给座位信息赋值，抱愧price，type，quantity
		print(result)
	}

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

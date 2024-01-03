package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"
	"go-zero-12306/common/comStruct"
	"go-zero-12306/common/constant"
	"go-zero-12306/common/tool"
	"strconv"
	"strings"
	"time"
)

type RegionTrainStationJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegionTrainStationJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegionTrainStationJobLogic {
	return &RegionTrainStationJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegionTrainStationJobLogic) RegionTrainStationJob(in *pb.RegionTrainStationJobReq) (*pb.RegionTrainStationJobResp, error) {
	// 1.查询所有region
	builder := l.svcCtx.TRegionModel.SelectBuilder()
	tRegions, err := l.svcCtx.TRegionModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return &pb.RegionTrainStationJobResp{}, err
	}
	tomorrowDateStr := tool.TomorrowDateStr()
	l.svcCtx.RedisClient.Set(l.ctx, "hello", "world", 15*24*time.Hour)
	fmt.Println(tomorrowDateStr)
	selectBuilder := l.svcCtx.TTrainStationRelationModel.SelectBuilder()
	// 2.遍历region，查询region下的所有站点
	for i := 0; i < len(tRegions); i++ {
		for j := 0; j < len(tRegions); j++ {
			if i != j {
				startRegin := tRegions[i].Name
				endRegion := tRegions[j].Name
				stationRelations, err := l.svcCtx.
					TTrainStationRelationModel.
					SelectSratr2End(l.ctx, selectBuilder, startRegin.String, endRegion.String)
				if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
					return &pb.RegionTrainStationJobResp{}, err
				}
				if len(stationRelations) == 0 {
					continue
				}
				var tuples = make(comStruct.TrainStationRelationRedis)
				for _, item := range stationRelations {
					key := strings.Join([]string{strconv.FormatInt(item.TrainId.Int64, 10), item.Departure.String, item.Arrival.String}, "_")
					jsonValue, _ := json.Marshal(&item)
					tuples[key] = string(jsonValue)
				}
				// 存入缓存 constant.REGION_TRAIN_STATION ，存15天
				//buildCacheKey := fmt.Sprintf(constant.REGION_TRAIN_STATION, startRegin.String, endRegion.String, tomorrowDateStr)
				buildCacheKey := fmt.Sprintf(constant.REGION_TRAIN_STATION, startRegin.String, endRegion.String)
				err = l.svcCtx.RedisClient.Set(l.ctx, buildCacheKey, tuples, constant.ADVANCE_TICKET_DAY).Err()
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return &pb.RegionTrainStationJobResp{}, nil
}

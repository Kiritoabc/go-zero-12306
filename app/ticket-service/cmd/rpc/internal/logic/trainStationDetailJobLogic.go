package logic

import (
	"context"
	"go-zero-12306/app/ticket-service/model/tTrain"
	"go-zero-12306/common/constant"
	"go-zero-12306/common/tool"
	"strconv"
	"strings"
	"time"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type TrainStationDetailJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTrainStationDetailJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrainStationDetailJobLogic {
	return &TrainStationDetailJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TrainStationDetailJobLogic) TrainStationDetailJob(in *pb.TrainStationDetailJobReq) (*pb.TrainStationDetailJobResp, error) {
	// todo: add your logic here and delete this linegt
	var currentPage int64 = 1
	var size int64 = 1000
	var dateTime time.Time
	if in.DateTime == "" {
		dateTime = tool.TomorrowDate()
	} else {
		inTime, _ := time.Parse(constant.Datetimeform1, in.DateTime)
		dateTime = inTime
	}
	tTrainBuilder := l.svcCtx.TTrainModel.SelectBuilder()
	for ; ; currentPage++ {
		listByPage, err := l.svcCtx.TTrainModel.FindPageListByPage(l.ctx, tTrainBuilder, size, currentPage, dateTime, "")
		if err != nil {
			return &pb.TrainStationDetailJobResp{}, err
		}
		if listByPage == nil || len(listByPage) == 0 {
			break
		}
		err = l.actualExecute(listByPage)
	}
	return &pb.TrainStationDetailJobResp{}, nil
}

func (l *TrainStationDetailJobLogic) actualExecute(trainDOPageRecords []*tTrain.TTrain) error {
	builder := l.svcCtx.TTrainStationRelationModel.SelectBuilder()
	redisClient := l.svcCtx.RedisClient
	for i := 0; i < len(trainDOPageRecords); i++ {
		list, err := l.svcCtx.TTrainStationRelationModel.SelectList(l.ctx, builder, strconv.FormatInt(trainDOPageRecords[i].Id, 10))
		if err != nil {
			return err
		}
		if len(list) == 0 {
			return nil
		}
		for _, item := range list {
			var actualCacheHashValue = map[string]string{}
			actualCacheHashValue["trainNumber"] = trainDOPageRecords[i].TrainNumber.String
			actualCacheHashValue["departureFlag"] = strconv.FormatInt(item.DepartureFlag.Int64, 10)
			actualCacheHashValue["arrivalFlag"] = strconv.FormatInt(item.ArrivalFlag.Int64, 10)
			actualCacheHashValue["departureTime"] = item.DepartureTime.Time.Format(constant.Timetemplate2)
			actualCacheHashValue["arrivalTime"] = item.ArrivalTime.Time.Format(constant.Timetemplate2)
			actualCacheHashValue["saleTime"] = trainDOPageRecords[i].SaleTime.Time.Format(constant.Timetemplate1)
			actualCacheHashValue["trainTag"] = trainDOPageRecords[i].TrainTag.String
			buildCacheKey := constant.TRAIN_STATION_DETAIL + strings.Join([]string{
				strconv.FormatInt(trainDOPageRecords[i].Id, 10),
				item.Departure.String,
				item.Arrival.String,
			}, "_")
			// 存入缓存
			err = redisClient.HMSet(l.ctx, buildCacheKey, actualCacheHashValue).Err()
			if err != nil {
				return err
			}
			redisClient.Expire(l.ctx, buildCacheKey, constant.ADVANCE_TICKET_DAY)
		}
	}
	return nil
}

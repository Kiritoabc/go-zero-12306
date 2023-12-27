package logic

import (
	"context"
	"strconv"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTrainStationQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTrainStationQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTrainStationQueryLogic {
	return &ListTrainStationQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListTrainStationQueryLogic) ListTrainStationQuery(in *pb.ListTrainStationQueryReq) (*pb.ListTrainStationQueryResp, error) {
	builder := l.svcCtx.TTrainStationModel.SelectBuilder()
	id64, err := strconv.ParseInt(in.TrainId, 10, 64)
	if err != nil {
		return &pb.ListTrainStationQueryResp{}, err
	}
	list, err := l.svcCtx.TTrainStationModel.SlectListByTrainId(l.ctx, builder, id64)
	if err != nil {
		return &pb.ListTrainStationQueryResp{}, err
	}
	var respList = []*pb.TrainStationQueryRespDTO{}
	for i := 0; i < len(list); i++ {
		var respDTO = &pb.TrainStationQueryRespDTO{}
		respDTO.Sequence = list[i].Sequence
		respDTO.Departure = list[i].Departure
		respDTO.ArrivalTime = list[i].ArrivalTime.Unix()
		respDTO.DepartureTime = list[i].DepartureTime.Unix()
		respDTO.StopoverTime = list[i].StopoverTime.Int64
		respList = append(respList, respDTO)
	}
	if err != nil {
		return &pb.ListTrainStationQueryResp{}, err
	}
	return &pb.ListTrainStationQueryResp{
		TrainStationQueryRespDTO: respList,
	}, nil
}

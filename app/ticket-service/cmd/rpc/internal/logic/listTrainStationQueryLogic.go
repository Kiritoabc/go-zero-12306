package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.ListTrainStationQueryResp{}, nil
}

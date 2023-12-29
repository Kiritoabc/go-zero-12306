package logic

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *ListRegionStationLogic) ListRegionStation(in *pb.ListRegionStationReq) (*pb.ListRegionStationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListRegionStationResp{}, nil
}

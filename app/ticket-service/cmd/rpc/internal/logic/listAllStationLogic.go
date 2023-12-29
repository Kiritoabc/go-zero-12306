package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.ListAllStationResp{}, nil
}

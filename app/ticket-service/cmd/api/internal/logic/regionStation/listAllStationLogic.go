package regionStation

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"go-zero-12306/app/ticket-service/cmd/api/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAllStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllStationLogic {
	return &ListAllStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAllStationLogic) ListAllStation(req *types.StationQueryReq) ([]*types.StationQueryRespDTO, error) {
	var resp []*types.StationQueryRespDTO
	listAllStation, err := l.svcCtx.TicketRpc.ListAllStation(l.ctx, &pb.ListAllStationReq{})
	if err != nil {
		return nil, err
	}
	copier.Copy(&resp, &listAllStation.StationQueryRespDTO)
	return resp, nil
}

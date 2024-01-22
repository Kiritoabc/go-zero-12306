package regionStation

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/ticket-service/cmd/api/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/types"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRegionStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRegionStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRegionStationLogic {
	return &ListRegionStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRegionStationLogic) ListRegionStation(req *types.RegionStationQueryReqDTO) (*types.RegionStationQueryResp, error) {
	var resp *types.RegionStationQueryResp
	listRegionStation, err := l.svcCtx.TicketRpc.ListRegionStation(l.ctx, &pb.ListRegionStationReq{
		QueryType: req.QuerryType,
		Name:      req.Name,
	})
	if err != nil {
		return resp, err
	}
	err = copier.Copy(&resp.List, &listRegionStation)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

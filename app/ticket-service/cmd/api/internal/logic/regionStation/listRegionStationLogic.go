package regionStation

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/types"

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

func (l *ListRegionStationLogic) ListRegionStation(req *types.RegionStationQueryReqDTO) (resp *types.RegionStationQueryResp, err error) {
	// todo: add your logic here and delete this line

	return
}

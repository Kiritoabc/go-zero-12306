package regionStation

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/types"

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

func (l *ListAllStationLogic) ListAllStation(req *types.StationQueryReq) (resp *types.StationQueryRespDTO, err error) {
	// todo: add your logic here and delete this line
	return
}

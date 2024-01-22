package trainStation

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/ticket-service/cmd/rpc/ticket"

	"go-zero-12306/app/ticket-service/cmd/api/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTrainStationQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTrainStationQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTrainStationQueryLogic {
	return &ListTrainStationQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTrainStationQueryLogic) ListTrainStationQuery(req *types.ListTrainStationQueryReq) (*types.TrainStationQueryResp, error) {
	data, err := l.svcCtx.TicketRpc.ListTrainStationQuery(l.ctx, &ticket.ListTrainStationQueryReq{
		TrainId: req.TrainId,
	})
	if err != nil {
		return nil, err
	}
	var list []types.TrainStationQueryRespDTO
	if data.TrainStationQueryRespDTO != nil {
		copier.Copy(&list, &data.TrainStationQueryRespDTO)
	}
	return &types.TrainStationQueryResp{
		List: list,
	}, nil
}

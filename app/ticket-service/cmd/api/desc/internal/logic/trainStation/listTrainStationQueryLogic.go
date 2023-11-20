package trainStation

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/types"

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

func (l *ListTrainStationQueryLogic) ListTrainStationQuery(req *types.ListTrainStationQueryReq) (resp *types.TrainStationQueryResp, err error) {
	// todo: add your logic here and delete this line

	return
}

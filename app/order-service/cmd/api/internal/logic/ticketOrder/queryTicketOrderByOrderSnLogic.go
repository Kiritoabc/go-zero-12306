package ticketOrder

import (
	"context"

	"go-zero-12306/app/order-service/cmd/api/internal/svc"
	"go-zero-12306/app/order-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTicketOrderByOrderSnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTicketOrderByOrderSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTicketOrderByOrderSnLogic {
	return &QueryTicketOrderByOrderSnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTicketOrderByOrderSnLogic) QueryTicketOrderByOrderSn(req *types.QueryTicketOrderByOrderSnReq) (resp *types.QueryTicketOrderByOrderSnResp, err error) {
	// todo: add your logic here and delete this line

	return
}

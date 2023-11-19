package ticketOrder

import (
	"context"

	"go-zero-12306/app/order-service/cmd/api/internal/svc"
	"go-zero-12306/app/order-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelTicketOrderReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelTicketOrderReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelTicketOrderReqLogic {
	return &CancelTicketOrderReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelTicketOrderReqLogic) CancelTicketOrderReq(req *types.CancelTicketOrderReq) (resp *types.CancelTicketOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}

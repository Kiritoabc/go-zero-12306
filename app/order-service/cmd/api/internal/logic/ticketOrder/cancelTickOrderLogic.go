package ticketOrder

import (
	"context"

	"go-zero-12306/app/order-service/cmd/api/internal/svc"
	"go-zero-12306/app/order-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelTickOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelTickOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelTickOrderLogic {
	return &CancelTickOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelTickOrderLogic) CancelTickOrder(req *types.CancelTicketOrderReq) (resp *types.CancelTicketOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}

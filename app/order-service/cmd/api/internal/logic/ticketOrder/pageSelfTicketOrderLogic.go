package ticketOrder

import (
	"context"

	"go-zero-12306/app/order-service/cmd/api/internal/svc"
	"go-zero-12306/app/order-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageSelfTicketOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageSelfTicketOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageSelfTicketOrderLogic {
	return &PageSelfTicketOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageSelfTicketOrderLogic) PageSelfTicketOrder(req *types.TicketOrderSelfPageQueryReq) (resp *types.TicketOrderSelfPageQueryResp, err error) {
	// todo: add your logic here and delete this line

	return
}

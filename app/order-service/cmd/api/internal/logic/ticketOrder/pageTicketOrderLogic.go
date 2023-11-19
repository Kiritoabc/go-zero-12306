package ticketOrder

import (
	"context"

	"go-zero-12306/app/order-service/cmd/api/internal/svc"
	"go-zero-12306/app/order-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageTicketOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageTicketOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageTicketOrderLogic {
	return &PageTicketOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageTicketOrderLogic) PageTicketOrder(req *types.PageTicketOrderReq) (resp *types.PageTicketOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}

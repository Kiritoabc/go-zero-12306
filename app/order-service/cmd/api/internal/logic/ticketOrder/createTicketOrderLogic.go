package ticketOrder

import (
	"context"

	"go-zero-12306/app/order-service/cmd/api/internal/svc"
	"go-zero-12306/app/order-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTicketOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTicketOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTicketOrderLogic {
	return &CreateTicketOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTicketOrderLogic) CreateTicketOrder(req *types.CreateTicketOrderReq) (resp *types.CreateTicketOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}

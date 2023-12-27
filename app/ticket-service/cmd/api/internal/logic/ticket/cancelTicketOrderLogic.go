package ticket

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelTicketOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelTicketOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelTicketOrderLogic {
	return &CancelTicketOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelTicketOrderLogic) CancelTicketOrder(req *types.CancelTicketOrderReqDTO) (resp *types.CancelTicketOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}

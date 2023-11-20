package ticket

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonTicketRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommonTicketRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonTicketRefundLogic {
	return &CommonTicketRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommonTicketRefundLogic) CommonTicketRefund(req *types.RefundTicketReqDTO) (resp *types.RefundTicketRespDTO, err error) {
	// todo: add your logic here and delete this line

	return
}

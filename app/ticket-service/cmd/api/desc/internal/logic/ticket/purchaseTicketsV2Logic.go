package ticket

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PurchaseTicketsV2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPurchaseTicketsV2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *PurchaseTicketsV2Logic {
	return &PurchaseTicketsV2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PurchaseTicketsV2Logic) PurchaseTicketsV2(req *types.PurchaseTicketReqDTO) (resp *types.TicketPurchaseRespDTO, err error) {
	// todo: add your logic here and delete this line

	return
}

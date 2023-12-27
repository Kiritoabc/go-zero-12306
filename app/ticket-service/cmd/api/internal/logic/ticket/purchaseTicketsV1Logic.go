package ticket

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PurchaseTicketsV1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPurchaseTicketsV1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *PurchaseTicketsV1Logic {
	return &PurchaseTicketsV1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PurchaseTicketsV1Logic) PurchaseTicketsV1(req *types.PurchaseTicketReqDTO) (resp *types.TicketPurchaseRespDTO, err error) {
	// todo: add your logic here and delete this line

	return
}

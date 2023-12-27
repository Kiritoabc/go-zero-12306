package tempSeat

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PurchaseTicketsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPurchaseTicketsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PurchaseTicketsLogic {
	return &PurchaseTicketsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PurchaseTicketsLogic) PurchaseTickets(req *types.PurchaseTicketsReq) (resp *types.PurchaseTicketsResp, err error) {
	// todo: add your logic here and delete this line

	return
}

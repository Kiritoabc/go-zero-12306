package logic

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PurchaseTicketsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPurchaseTicketsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PurchaseTicketsLogic {
	return &PurchaseTicketsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PurchaseTicketsLogic) PurchaseTickets(in *pb.PurchaseTicketsReq) (*pb.PurchaseTicketsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PurchaseTicketsResp{}, nil
}

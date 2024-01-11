package logic

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PurchaseTicketsV2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPurchaseTicketsV2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *PurchaseTicketsV2Logic {
	return &PurchaseTicketsV2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PurchaseTicketsV2Logic) PurchaseTicketsV2(in *pb.PurchaseTicketsReq) (*pb.PurchaseTicketsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PurchaseTicketsResp{}, nil
}

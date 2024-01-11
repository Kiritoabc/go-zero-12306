package logic

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelTicketOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelTicketOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelTicketOrderLogic {
	return &CancelTicketOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelTicketOrderLogic) CancelTicketOrder(in *pb.CancelTicketOrderReq) (*pb.CancelTicketOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CancelTicketOrderResp{}, nil
}

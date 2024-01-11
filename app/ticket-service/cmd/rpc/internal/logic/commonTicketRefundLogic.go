package logic

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonTicketRefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommonTicketRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonTicketRefundLogic {
	return &CommonTicketRefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommonTicketRefundLogic) CommonTicketRefund(in *pb.CommonTicketRefundReq) (*pb.CommonTicketRefundResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CommonTicketRefundResp{}, nil
}

package logic

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageListTicketQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageListTicketQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageListTicketQueryLogic {
	return &PageListTicketQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TODO: TicketControllerRpc
func (l *PageListTicketQueryLogic) PageListTicketQuery(in *pb.PageListTicketQueryReq) (*pb.PageListTicketQueryResp, error) {
	// todo: 责任链模式？？emmmmm

	return &pb.PageListTicketQueryResp{}, nil
}

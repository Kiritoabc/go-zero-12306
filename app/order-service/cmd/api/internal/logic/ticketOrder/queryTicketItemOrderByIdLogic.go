package ticketOrder

import (
	"context"

	"go-zero-12306/app/order-service/cmd/api/internal/svc"
	"go-zero-12306/app/order-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTicketItemOrderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTicketItemOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTicketItemOrderByIdLogic {
	return &QueryTicketItemOrderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTicketItemOrderByIdLogic) QueryTicketItemOrderById(req *types.QueryTicketItemOrderByIdReq) (resp *types.QueryTicketItemOrderByIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}

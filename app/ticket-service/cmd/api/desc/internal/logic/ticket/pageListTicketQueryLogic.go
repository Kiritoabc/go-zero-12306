package ticket

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageListTicketQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageListTicketQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageListTicketQueryLogic {
	return &PageListTicketQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageListTicketQueryLogic) PageListTicketQuery(req *types.TicketPageQueryReqDTO) (resp *types.TicketPageQueryRespDTO, err error) {
	// todo: add your logic here and delete this line

	return
}

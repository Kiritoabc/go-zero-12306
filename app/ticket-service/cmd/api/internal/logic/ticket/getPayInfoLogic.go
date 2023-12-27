package ticket

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/api/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPayInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPayInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPayInfoLogic {
	return &GetPayInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPayInfoLogic) GetPayInfo(req *types.GetPayInfoReq) (resp *types.PayInfoRespDTO, err error) {
	// todo: add your logic here and delete this line

	return
}

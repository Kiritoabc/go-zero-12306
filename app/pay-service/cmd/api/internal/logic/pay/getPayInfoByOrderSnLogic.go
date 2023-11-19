package pay

import (
	"context"

	"go-zero-12306/app/pay-service/cmd/api/internal/svc"
	"go-zero-12306/app/pay-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPayInfoByOrderSnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPayInfoByOrderSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPayInfoByOrderSnLogic {
	return &GetPayInfoByOrderSnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPayInfoByOrderSnLogic) GetPayInfoByOrderSn(req *types.GetPayInfoByOrderSnReq) (resp *types.GetPayInfoByOrderSnResp, err error) {
	// todo: add your logic here and delete this line

	return
}

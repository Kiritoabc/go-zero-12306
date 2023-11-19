package payCallback

import (
	"context"

	"go-zero-12306/app/pay-service/cmd/api/internal/svc"
	"go-zero-12306/app/pay-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackAlipayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallbackAlipayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackAlipayLogic {
	return &CallbackAlipayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallbackAlipayLogic) CallbackAlipay(req *types.CallbackAlipayReq) (resp *types.CallbackAlipayResp, err error) {
	// todo: add your logic here and delete this line

	return
}

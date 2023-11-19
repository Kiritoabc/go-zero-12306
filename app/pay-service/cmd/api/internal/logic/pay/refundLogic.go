package pay

import (
	"context"

	"go-zero-12306/app/pay-service/cmd/api/internal/svc"
	"go-zero-12306/app/pay-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefundLogic {
	return &RefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefundLogic) Refund(req *types.RefundReq) (resp *types.RefundResp, err error) {
	// todo: add your logic here and delete this line

	return
}

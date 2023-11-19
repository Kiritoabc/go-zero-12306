package refund

import (
	"context"

	"go-zero-12306/app/pay-service/cmd/api/internal/svc"
	"go-zero-12306/app/pay-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommonRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonRefundLogic {
	return &CommonRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommonRefundLogic) CommonRefund(req *types.RefundReq) (resp *types.RefundResp, err error) {
	// todo: add your logic here and delete this line

	return
}

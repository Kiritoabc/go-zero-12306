package pay

import (
	"context"

	"go-zero-12306/app/pay-service/cmd/api/internal/svc"
	"go-zero-12306/app/pay-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPayInfoByPaySnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPayInfoByPaySnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPayInfoByPaySnLogic {
	return &GetPayInfoByPaySnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPayInfoByPaySnLogic) GetPayInfoByPaySn(req *types.GetPayInfoByPaySnReq) (resp *types.GetPayInfoByPaySnResp, err error) {
	// todo: add your logic here and delete this line

	return
}

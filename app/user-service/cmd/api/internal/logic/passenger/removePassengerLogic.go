package passenger

import (
	"context"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePassengerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovePassengerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePassengerLogic {
	return &RemovePassengerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovePassengerLogic) RemovePassenger(req *types.RemovePassengerReq) (resp *types.RemovePassengerResp, err error) {
	// todo: add your logic here and delete this line

	return
}

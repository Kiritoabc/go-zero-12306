package passenger

import (
	"context"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePassengerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePassengerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePassengerLogic {
	return &UpdatePassengerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePassengerLogic) UpdatePassenger(req *types.UpdatePassengerReq) (resp *types.UpdatePassengerResp, err error) {
	// todo: add your logic here and delete this line

	return
}

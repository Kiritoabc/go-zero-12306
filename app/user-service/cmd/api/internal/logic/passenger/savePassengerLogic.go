package passenger

import (
	"context"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SavePassengerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSavePassengerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SavePassengerLogic {
	return &SavePassengerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SavePassengerLogic) SavePassenger(req *types.SavePassengerReq) (resp *types.SavePassengerResp, err error) {
	// todo: add your logic here and delete this line

	return
}

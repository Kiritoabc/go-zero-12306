package passenger

import (
	"context"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPassengerQueryByUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPassengerQueryByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPassengerQueryByUsernameLogic {
	return &ListPassengerQueryByUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPassengerQueryByUsernameLogic) ListPassengerQueryByUsername(req *types.PassengerReq) (resp *types.PassengerRespDTO, err error) {
	// todo: add your logic here and delete this line

	return
}

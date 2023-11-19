package passenger

import (
	"context"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPassengerQueryByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPassengerQueryByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPassengerQueryByIdsLogic {
	return &ListPassengerQueryByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPassengerQueryByIdsLogic) ListPassengerQueryByIds(req *types.PassengerQueryReq) (resp *types.PassengerActualRespDTO, err error) {
	// todo: add your logic here and delete this line

	return
}

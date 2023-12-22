package passenger

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

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

func (l *ListPassengerQueryByIdsLogic) ListPassengerQueryByIds(req *types.PassengerQueryReq) (*types.PassengerQueryResp, error) {
	listResp, err := l.svcCtx.UserRpc.ListPassengerQueryByUsername(l.ctx, &pb.ListPassengerQueryByUsernameReq{
		Username: req.UserName,
	})
	if err != nil {
		return nil, err
	}
	var resp types.PassengerQueryResp
	err = copier.Copy(&resp, &listResp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

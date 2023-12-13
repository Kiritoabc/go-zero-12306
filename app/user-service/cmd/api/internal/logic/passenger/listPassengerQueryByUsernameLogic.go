package passenger

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

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

func (l *ListPassengerQueryByUsernameLogic) ListPassengerQueryByUsername(req *types.PassengerReq) (*types.PassengerResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.UserRpc.ListPassengerQueryByUsername(l.ctx, &pb.ListPassengerQueryByUsernameReq{
		Username: req.UserName,
	})
	if err != nil {
		return nil, err
	}
	var resp types.PassengerResp
	_ = copier.Copy(&resp, list)
	return &resp, nil
}

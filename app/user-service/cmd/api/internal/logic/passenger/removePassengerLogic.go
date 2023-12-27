package passenger

import (
	"context"
	"go-zero-12306/app/user-service/cmd/rpc/user"
	"strconv"

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

func (l *RemovePassengerLogic) RemovePassenger(req *types.RemovePassengerReq) (*types.RemovePassengerResp, error) {
	parseInt, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserRpc.RemovePassenger(l.ctx, &user.RemovePassengerReq{
		Id: parseInt,
	})
	if err != nil {
		return &types.RemovePassengerResp{}, err
	}
	return &types.RemovePassengerResp{}, nil
}

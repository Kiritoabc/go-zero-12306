package userInfo

import (
	"context"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HasUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHasUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HasUsernameLogic {
	return &HasUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HasUsernameLogic) HasUsername(req *types.UserHasUsernameReq) (*types.UserHasUsernameResp, error) {
	username, err := l.svcCtx.UserRpc.HasUsername(l.ctx, &pb.HasUsernameReq{
		Username: req.UserName,
	})
	if err != nil {
		return nil, err
	}
	var resp types.UserHasUsernameResp
	resp.Bool = username.Has
	return &resp, err
}

package userLogin

import (
	"context"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (*types.LogoutResp, error) {
	// 调用rpc
	_, err := l.svcCtx.UserRpc.Logout(l.ctx, &pb.LogoutReq{
		AccessToken: req.AccessToken,
	})
	if err != nil {
		return nil, err
	}
	return &types.LogoutResp{}, nil
}

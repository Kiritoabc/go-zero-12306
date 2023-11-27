package userInfo

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserByUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserByUsernameLogic {
	return &QueryUserByUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserByUsernameLogic) QueryUserByUsername(req *types.QueryUserReq) (*types.QueryUserResp, error) {
	// todo: add your logic here and delete this line
	UserDO, err := l.svcCtx.UserRpc.QueryUserByUsername(l.ctx, &pb.UserNameReq{Username: req.UserName})
	if err != nil {
		return nil, err
	}
	var resp types.QueryUserResp
	_ = copier.Copy(&resp, UserDO)
	return &resp, nil
}

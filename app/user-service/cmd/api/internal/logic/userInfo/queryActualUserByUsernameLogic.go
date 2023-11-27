package userInfo

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryActualUserByUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryActualUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryActualUserByUsernameLogic {
	return &QueryActualUserByUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryActualUserByUsernameLogic) QueryActualUserByUsername(req *types.UserQueryActualReq) (*types.UserQueryActualResp, error) {
	// todo: add your logic here and delete this line
	UserDO, err := l.svcCtx.UserRpc.QueryUserByUsername(l.ctx, &pb.UserNameReq{Username: req.UserName})
	if err != nil {
		return nil, err
	}
	var resp types.UserQueryActualResp
	_ = copier.Copy(&resp, &UserDO)
	return &resp, nil
}

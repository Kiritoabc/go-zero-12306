package userInfo

import (
	"context"

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

func (l *QueryUserByUsernameLogic) QueryUserByUsername(req *types.QueryUserReq) (resp *types.QueryUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}

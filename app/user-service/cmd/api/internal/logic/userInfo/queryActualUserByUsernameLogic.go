package userInfo

import (
	"context"

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

func (l *QueryActualUserByUsernameLogic) QueryActualUserByUsername(req *types.UserQueryActualReq) (resp *types.UserQueryActualResp, err error) {
	// todo: add your logic here and delete this line

	return
}

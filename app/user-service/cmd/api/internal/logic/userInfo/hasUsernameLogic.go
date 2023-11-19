package userInfo

import (
	"context"

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

func (l *HasUsernameLogic) HasUsername(req *types.UserHasUsernameReq) (resp *types.UserHasUsernameResp, err error) {
	// todo: add your logic here and delete this line

	return
}

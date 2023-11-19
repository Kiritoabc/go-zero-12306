package userInfo

import (
	"context"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletionLogic {
	return &DeletionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletionLogic) Deletion(req *types.UserDeleteReq) (resp *types.UserDeleteResp, err error) {
	// todo: add your logic here and delete this line

	return
}

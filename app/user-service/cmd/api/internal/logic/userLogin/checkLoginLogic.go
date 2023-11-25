package userLogin

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLoginLogic {
	return &CheckLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLoginLogic) CheckLogin(req *types.CheckLoginReq) (*types.CheckLoginResp, error) {
	// 调用rpc服务
	rpcResp, err := l.svcCtx.UserRpc.CheckLogin(l.ctx, &pb.CheckLoginReq{
		AccessToken: req.AccessToken,
	})
	if err != nil {
		return nil, err
	}
	var checkResp types.CheckLoginResp
	err = copier.Copy(&checkResp, &rpcResp)
	if err != nil {
		return nil, err
	}
	return &checkResp, nil
}

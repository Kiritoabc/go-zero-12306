package userLogin

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/user-service/cmd/rpc/user"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// 调用rpc服务
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		UsernameOrMailOrPhone: req.UserNameOrMailOrPhone,
		Password:              req.Password,
	})
	if err != nil {
		return nil, err
	}

	var resp types.LoginResp
	_ = copier.Copy(&resp, loginResp)
	return &resp, nil
}

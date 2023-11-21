package userInfo

import (
	"context"
	"go-zero-12306/app/user-service/cmd/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	// todo: add your logic here and delete this line
	// 调用rpc服务
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		UserName:    req.UserName,
		Password:    req.Password,
		RealName:    req.RealName,
		IdType:      req.IdType,
		IdCard:      req.IdCard,
		Phone:       req.Phone,
		Mail:        req.Mail,
		UserType:    req.UserType,
		VerifyState: req.VerifyState,
		PostCode:    req.PostCode,
		Address:     req.Address,
		Region:      req.Region,
		Telephone:   req.Telephone,
	})
	if err != nil {
		return nil, err
	}
	// copier
	var resp *types.RegisterResp
	_ = copier.Copy(&resp, &registerResp)
	return resp, nil
}

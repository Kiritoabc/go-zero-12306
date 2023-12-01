package userInfo

import (
	"context"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		Id:       req.Id,
		Username: req.UserName,
		Mail:     req.Mail,
		UserType: req.UserType,
		PostCode: req.PostCode,
		Address:  req.Address,
	})
	if err != nil {
		return nil, err
	}
	return
}

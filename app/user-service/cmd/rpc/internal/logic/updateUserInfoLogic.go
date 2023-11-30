package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"
	"go-zero-12306/app/user-service/model/tUser"
	"go-zero-12306/app/user-service/model/tUserMail"
	"go-zero-12306/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	// todo: add your logic here and delete this line
	//首先查询到
	user, err := l.svcCtx.User0Model.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询user失败:%v", err)
	}
	var updateUser tUser.TUser0
	_ = copier.Copy(&updateUser, &in)
	// 更新 user
	err = l.svcCtx.User0Model.Update(l.ctx, &updateUser)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user更新失败:%v", err)
	}
	// 更新mail，首先删除mail，再更新mail
	_ = l.svcCtx.UserMail0Model.DeleteByMail(l.ctx, user.Mail)
	var updateMail tUserMail.TUserMail0
	_ = copier.Copy(&updateMail, in)
	_, _ = l.svcCtx.UserMail0Model.Insert(l.ctx, nil, &updateMail)
	return &pb.UpdateUserInfoResp{}, nil
}

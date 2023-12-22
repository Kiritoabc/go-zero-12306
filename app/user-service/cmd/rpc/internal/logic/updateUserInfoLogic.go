package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	// 开启事务
	if err := l.svcCtx.User0Model.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user, err := l.svcCtx.User0Model.FindOneByUsername(l.ctx, session, in.Username)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询user失败:%v", err)
		}
		var updateUser tUser.TUser0
		_ = copier.Copy(&updateUser, &in)
		// 更新user
		err = l.svcCtx.User0Model.Update(l.ctx, session, &updateUser)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user更新失败:%v", err)
		}
		// 更新mail，首先删除mail，再更新mail
		err = l.svcCtx.UserMail0Model.DeleteByMail(l.ctx, session, user.Mail)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "email删除失败:%v", err)
		}
		var updateMail tUserMail.TUserMail0
		_ = copier.Copy(&updateMail, in)
		_, err = l.svcCtx.UserMail0Model.Insert(l.ctx, nil, &updateMail)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "email更新失败:%v", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateUserInfoResp{}, nil
}

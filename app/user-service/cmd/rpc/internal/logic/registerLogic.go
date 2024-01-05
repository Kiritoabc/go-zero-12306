package logic

import (
	"context"
	godisson "github.com/cheerego/go-redisson"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"
	"go-zero-12306/app/user-service/model/tUser"
	"go-zero-12306/app/user-service/model/tUserMail"
	"go-zero-12306/app/user-service/model/tUserPhone"
	"go-zero-12306/common/constant"
	"go-zero-12306/common/tool"
	"go-zero-12306/common/xerr"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 开启事务

	// 1.redisson上锁
	g := godisson.NewGodisson(l.svcCtx.RedisClient, godisson.WithWatchDogTimeout(30*time.Second))
	// lock with watchdog without retry
	lock := g.NewRLock(constant.LOCK_USER_REGISTER + in.UserName)
	err := lock.Lock()
	if err == godisson.ErrLockNotObtained {
		return &pb.RegisterResp{}, errors.Wrapf(xerr.NewErrCode(xerr.HAS_USERNAME_NOTNULL), "用户名已存在:%v,用户名:%+v", err, in.UserName)
	} else if err != nil {
		return &pb.RegisterResp{}, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "redis获取锁出错")
	}
	// 7.解锁
	defer lock.Unlock()
	if err := l.svcCtx.User0Model.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 0.参数检查(暂时不做考虑)
		// 2.插入用户，判断用户名是否重复
		user := new(tUser.TUser0)
		_ = copier.Copy(&user, &in)
		user.Password = tool.Md5ByString(in.Password)
		_, err = l.svcCtx.User0Model.Insert(ctx, session, user) // 插入用户
		// 这里应该判断一下是否是用户名重复
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
		}
		// 3.插入手机号，检查手机号是否重复
		userPhone0 := new(tUserPhone.TUserPhone0)
		_ = copier.Copy(&userPhone0, &in)
		_, err = l.svcCtx.UserPhone0Model.Insert(ctx, session, userPhone0)
		// 判断是手机号重复的错误
		if err != nil {
			return err
		}
		// 4.插入邮箱，检查邮箱是否重复
		userMail0 := new(tUserMail.TUserMail0)
		_ = copier.Copy(&userMail0, &in)
		_, err = l.svcCtx.UserMail0Model.Insert(ctx, session, userMail0)
		// 判断邮箱是否重复
		if err != nil {
			return err
		}
		// 5.注册账号的时候，把在注销账号里面的清除掉
		err = l.svcCtx.UserReuseModel.DeleteByUserName(ctx, session, in.UserName)
		if err != nil {
			return err
		}
		err = l.svcCtx.RedisClient1.Set(in.UserName, "OK")
		if err != nil {
			return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "redis存储出错")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		UserName: in.UserName,
		RealName: in.RealName,
		Phone:    in.Phone,
	}, nil
}

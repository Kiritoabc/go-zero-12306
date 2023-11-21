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
	"go-zero-12306/common/constant"
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
	// todo: 注册服务的逻辑
	// 0.参数检查(暂时不做考虑)
	// 1.redisson上锁
	g := godisson.NewGodisson(l.svcCtx.RedisClient, godisson.WithWatchDogTimeout(30*time.Second))
	// lock with watchdog without retry
	lock := g.NewRLock(constant.LOCK_USER_REGISTER + in.UserName)
	err := lock.Lock()
	if err == godisson.ErrLockNotObtained {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.HAS_USERNAME_NOTNULL), "用户名已存在:%v,用户名:%+v", err, in.UserName)
	} else if err != nil {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "redis获取锁出错")
	}
	// 开启事务
	if err := l.svcCtx.User0Model.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 2.插入用户，判断用户名是否重复
		user := new(tUser.TUser0)
		_ = copier.Copy(user, in)
		// 3.插入手机号，检查手机号是否重复
		// 4.插入邮箱，检查邮箱是否重复
		// 5.注册账号的时候，把在注销账号里面的清除掉
		// 6.设置布隆过滤器
		return nil
	}); err != nil {
		return nil, err
	}
	// 7.解锁
	defer lock.Unlock()
	return nil, nil
}

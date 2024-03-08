package logic

import (
	"context"
	"fmt"
	godisson "github.com/cheerego/go-redisson"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/user-service/model/tUserDeletion"
	"go-zero-12306/app/user-service/model/tUserReuse"
	"go-zero-12306/common/constant"
	"go-zero-12306/common/globalkey"
	"go-zero-12306/common/xerr"
	"time"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletionLogic {
	return &DeletionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletionLogic) Deletion(in *pb.DeletionReq) (*pb.DeletionResp, error) {
	userId := in.GetId()
	// 1.判断注销的用户是否是本身
	user, err := l.svcCtx.User0Model.FindOneByUsername(l.ctx, nil, in.Username)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DB error: %v", err.Error())
	}
	if userId != user.Id {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.NAME_ERROR), "username is error: %v", in.Username)
	}
	// 锁放在外面
	g := godisson.NewGodisson(l.svcCtx.RedisClient, godisson.WithWatchDogTimeout(30*time.Second))
	lock := g.NewRLock(constant.LOCK_USER_REGISTER + in.Username)
	// 2.上锁
	err = lock.Lock()
	if err == godisson.ErrLockNotObtained {
		return &pb.DeletionResp{}, errors.Wrapf(xerr.NewErrCode(xerr.HAS_USERNAME_NOTNULL), "用户名已存在:%v,用户名:%+v", err, in.Username)
	} else if err != nil {
		return &pb.DeletionResp{}, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "redis获取锁出错")
	}
	// 解锁
	defer func() {
		_, err := lock.Unlock()
		if err != nil {
			fmt.Printf("解锁出错:%v", err)
		}
	}()

	// 开启事务
	if err = l.svcCtx.User0Model.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {

		// 3.插入到注销表
		var userDeletion tUserDeletion.TUserDeletion
		_ = copier.Copy(&userDeletion, user)
		l.svcCtx.UserDeletionModel.InsertWithSession(context, session, &userDeletion)
		// 4. 删除user
		l.svcCtx.User0Model.DeleteWithSession(context, session, user.Id)
		// 5. 删除phone
		l.svcCtx.UserPhone0Model.DeleteWithSession(context, session, user.Phone)
		if user.Mail != "" {
			// 删除mail
			l.svcCtx.UserMail0Model.DeleteByMail(context, session, user.Mail)
		}
		// 6.删除redis中的session
		userCacheTokenKey := fmt.Sprintf(globalkey.CacheUserTokenKey, userId)
		l.svcCtx.RedisClient.Del(l.ctx, userCacheTokenKey)
		// 7.Resume插入
		l.svcCtx.UserReuseModel.Insert(l.ctx, session, &tUserReuse.TUserReuse{
			Username: user.Username,
		})
		// todo: 8.布隆过滤器删除
		return nil
	}); err != nil {
		return &pb.DeletionResp{}, err
	}

	return &pb.DeletionResp{}, nil
}

package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"
	"go-zero-12306/app/user-service/cmd/rpc/user"
	"go-zero-12306/common/globalkey"
	"go-zero-12306/common/tool"
	"go-zero-12306/common/xerr"
	"strconv"
	"time"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 1. 判断usernameOrMailOrPhone是否是mail
	usernameOrMailOrPhone := in.UsernameOrMailOrPhone
	var mailFlag bool
	for i := range usernameOrMailOrPhone {
		if usernameOrMailOrPhone[i] == '@' {
			mailFlag = true
			break
		}
	}
	var username string
	if mailFlag {
		mailDO, err := l.svcCtx.UserMail0Model.FindOneByMail(l.ctx, usernameOrMailOrPhone)
		if err != nil {
			return &pb.LoginResp{}, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "数据库查找失败")
		}
		if mailDO == nil {
			return &pb.LoginResp{}, errors.Wrapf(xerr.NewErrCode(xerr.LOGIN_MAIL_NOT_EXIST), "用户名/手机号/邮箱不存在")
		}
		username = mailDO.Username
	} else {
		phoneDO, err := l.svcCtx.UserPhone0Model.FindOneByPhone(l.ctx, usernameOrMailOrPhone)

		if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
			return &pb.LoginResp{}, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "数据库查找失败")
		}
		if phoneDO != nil {
			username = phoneDO.Username
		}
	}
	// 2. 找出username
	if username == "" {
		username = usernameOrMailOrPhone
	}
	// 3. 查询
	userDO, err := l.svcCtx.User0Model.FindOneByUsername(l.ctx, nil, username)
	if userDO == nil {
		return &pb.LoginResp{}, errors.Wrapf(xerr.NewErrCode(xerr.LOGIN_MAIL_NOT_EXIST), "用户名/手机号/邮箱不存在")
	}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查找失败")
	}

	// 判断密码
	if tool.Md5ByString(in.Password) != userDO.Password {
		return &pb.LoginResp{}, errors.Wrapf(xerr.NewErrCode(xerr.LOGIN_MAIL_NOT_EXIST), "密码错误")
	}
	userId := strconv.FormatInt(userDO.Id, 10)
	// 4. 生成 accessToken
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: userDO.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}
	loginresp := &pb.LoginResp{
		UserId:      userId,
		Username:    userDO.Username,
		RealName:    userDO.RealName,
		AccessToken: tokenResp.AccessToken,
	}
	// 5. 将accessToken存入redis缓存
	loginRespJson, err := json.Marshal(&loginresp)
	if err != nil {
		return &pb.LoginResp{}, err
	}
	// todo ；暂时采用10分钟的测试
	usderIdKey, _ := strconv.ParseInt(userId, 10, 64)
	userCacheTokenKey := fmt.Sprintf(globalkey.CacheUserTokenKey, usderIdKey)
	err = l.svcCtx.RedisClient.SetEX(l.ctx, userCacheTokenKey, loginRespJson, 10*time.Minute).Err()
	err = l.svcCtx.RedisClient.SetEX(l.ctx, tokenResp.AccessToken, loginRespJson, 10*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	return loginresp, nil
}

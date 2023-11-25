package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"go-zero-12306/common/xerr"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLoginLogic {
	return &CheckLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLoginLogic) CheckLogin(in *pb.CheckLoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	// 获取到token
	accessToken := in.AccessToken
	// 从redis中判断用户是否登录
	stringCmd := l.svcCtx.RedisClient.Get(l.ctx, accessToken)
	bytes := []byte(stringCmd.Val())
	loginResp := &pb.LoginResp{}
	err := json.Unmarshal(bytes, &loginResp)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "json序列化失败")
	}
	return loginResp, nil
}

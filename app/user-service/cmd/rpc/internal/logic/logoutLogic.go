package logic

import (
	"context"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *pb.LogoutReq) (*pb.VoidResp, error) {
	// 1.获取到token
	accessToken := in.AccessToken
	// 2. 从redis中删除token
	l.svcCtx.RedisClient.Del(l.ctx, accessToken)
	return &pb.VoidResp{}, nil
}

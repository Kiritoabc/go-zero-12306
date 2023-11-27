package logic

import (
	"context"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HasUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHasUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HasUsernameLogic {
	return &HasUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HasUsernameLogic) HasUsername(in *pb.HasUsernameReq) (*pb.HasUsernameResp, error) {
	// 获取参数
	username := in.Username
	// 查询
	hasUsername, err := l.svcCtx.User0Model.FindOneByUsername(l.ctx, username)
	if err != nil {
		return nil, err
	}
	if hasUsername == nil {
		return &pb.HasUsernameResp{
			Has: false,
		}, nil
	}
	return &pb.HasUsernameResp{
		Has: true,
	}, nil
}

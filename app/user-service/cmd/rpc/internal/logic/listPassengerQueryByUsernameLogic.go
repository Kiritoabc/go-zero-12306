package logic

import (
	"context"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPassengerQueryByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPassengerQueryByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPassengerQueryByUsernameLogic {
	return &ListPassengerQueryByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPassengerQueryByUsernameLogic) ListPassengerQueryByUsername(in *pb.ListPassengerQueryByUsernameReq) (*pb.ListPassengerQueryByUsernameResp, error) {
	// todo: add your logic here and delete this line
	userName := in.Username

	return &pb.ListPassengerQueryByUsernameResp{}, nil
}

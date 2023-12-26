package logic

import (
	"context"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePassengerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemovePassengerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePassengerLogic {
	return &RemovePassengerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemovePassengerLogic) RemovePassenger(in *pb.RemovePassengerReq) (*pb.SavePassengerReq, error) {
	// todo: add your logic here and delete this line

	return &pb.SavePassengerReq{}, nil
}

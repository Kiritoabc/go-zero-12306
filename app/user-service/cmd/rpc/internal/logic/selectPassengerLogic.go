package logic

import (
	"context"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectPassengerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectPassengerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPassengerLogic {
	return &SelectPassengerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectPassengerLogic) SelectPassenger(in *pb.SelectPassengerReq) (*pb.PassengerDO, error) {
	// todo: add your logic here and delete this line

	return &pb.PassengerDO{}, nil
}

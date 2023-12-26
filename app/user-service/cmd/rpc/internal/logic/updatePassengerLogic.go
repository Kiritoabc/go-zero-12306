package logic

import (
	"context"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePassengerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePassengerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePassengerLogic {
	return &UpdatePassengerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePassengerLogic) UpdatePassenger(in *pb.SavePassengerReq) (*pb.SavePassengerResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SavePassengerResp{}, nil
}

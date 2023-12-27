package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	// 判断用户是否存在
	passenger, err := l.svcCtx.Passenger0Model.FindOne(l.ctx, in.Id)
	if err != nil {
		return &pb.SavePassengerReq{}, err
	}
	if err := l.svcCtx.Passenger0Model.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		err = l.svcCtx.Passenger0Model.Delete(l.ctx, session, in.Id, passenger.Username)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return &pb.SavePassengerReq{}, err
	}
	return &pb.SavePassengerReq{}, nil
}

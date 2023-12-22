package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-12306/app/user-service/model/tPassenger"

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
	userName := in.Username
	//1. 更具用户名查询
	// 使用了缓存
	var passengerList []*tPassenger.TPassenger0
	passengerList, err := l.svcCtx.Passenger0Model.GetPassengerByUserName(l.ctx, userName)
	if err != nil {
		return nil, err
	}
	var list []*pb.PassengerRespDTO
	// 赋值给resp
	if passengerList != nil {
		err = copier.Copy(&list, &passengerList)
	}
	return &pb.ListPassengerQueryByUsernameResp{
		List: list,
	}, nil
}

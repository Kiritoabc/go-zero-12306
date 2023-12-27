package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"
	"go-zero-12306/app/user-service/model/tPassenger"
	"strconv"
	"time"

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
	id64, _ := strconv.ParseInt(in.Id, 10, 64)
	passenger0, err := l.svcCtx.Passenger0Model.FindOne(l.ctx, id64)
	if err != nil {
		return nil, err
	}
	// 给passengerDO赋值
	var passengerDO = pb.PassengerDO{}
	err = copier.Copy(&passengerDO, &in)
	if err != nil {
		return &pb.SavePassengerResp{}, err
	}
	// 赋值username，VerifyStatus(0未审核,1已审核)
	passengerDO.Username = passenger0.Username
	// 开启事务
	if err = l.svcCtx.Passenger0Model.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var data = &tPassenger.TPassenger0{}
		err := copier.Copy(&data, &passengerDO)
		if err != nil {
			return err
		}
		data.UpdateTime = time.Now()
		data.Id = passenger0.Id
		err = l.svcCtx.Passenger0Model.UpdateWithTran(ctx, session, data)
		if err != nil {
			fmt.Println("新增乘车人失败")
			return err
		}
		return nil
	}); err != nil {
		return &pb.SavePassengerResp{}, err
	}
	return &pb.SavePassengerResp{}, nil
}

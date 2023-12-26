package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/user-service/model/tPassenger"
	"strconv"
	"time"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SavePassengerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSavePassengerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SavePassengerLogic {
	return &SavePassengerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SavePassengerLogic) SavePassenger(in *pb.SavePassengerReq) (*pb.SavePassengerResp, error) {
	// 首先从redis中取出accession
	id64, err := strconv.ParseInt(in.Id, 10, 64)
	user0, err := l.svcCtx.User0Model.FindOne(l.ctx, id64)
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
	passengerDO.Username = user0.Username
	passengerDO.CreateDate = time.Now().Unix()
	passengerDO.VerifyStatus = 1
	// 开启事务
	if err = l.svcCtx.Passenger0Model.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var data = &tPassenger.TPassenger0{}
		err := copier.Copy(&data, &passengerDO)
		if err != nil {
			return err
		}
		data.CreateDate = time.Now()
		data.Id = id64
		err = l.svcCtx.Passenger0Model.InsertWithTran(ctx, session, data)
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

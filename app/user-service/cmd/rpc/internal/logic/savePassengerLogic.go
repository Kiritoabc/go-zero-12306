package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"go-zero-12306/common/ctxdata"
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
	// todo: add your logic here and delete this line
	// 首先从redis中取出accession
	accession := ctxdata.GetAccessionFromCtx(l.ctx)
	LoginRespJson, err := l.svcCtx.RedisClient1.Get(accession)
	if err != nil {
		return &pb.SavePassengerResp{}, err
	}
	var loginResp *pb.LoginResp
	err = json.Unmarshal([]byte(LoginRespJson), &loginResp)
	if err != nil {
		return &pb.SavePassengerResp{}, err
	}
	// 给passengerDO赋值
	var passengerDO *pb.PassengerDO
	err = copier.Copy(&passengerDO, &in)
	if err != nil {
		return &pb.SavePassengerResp{}, err
	}
	// 赋值username
	passengerDO.Username = loginResp.Username
	passengerDO.CreateTime = time.Now().Unix()
	return &pb.SavePassengerResp{}, nil
}

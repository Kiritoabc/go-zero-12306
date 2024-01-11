package logic

import (
	"context"

	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPayInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPayInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPayInfoLogic {
	return &GetPayInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPayInfoLogic) GetPayInfo(in *pb.GetPayInfoReq) (*pb.GetPayInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPayInfoResp{}, nil
}

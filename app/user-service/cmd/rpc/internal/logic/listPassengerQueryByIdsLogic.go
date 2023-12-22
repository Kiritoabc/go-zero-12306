package logic

import (
	"context"

	"go-zero-12306/app/user-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/user-service/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPassengerQueryByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPassengerQueryByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPassengerQueryByIdsLogic {
	return &ListPassengerQueryByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPassengerQueryByIdsLogic) ListPassengerQueryByIds(in *pb.ListPassengerQueryByIdsReq) (*pb.ListPassengerQueryByIdsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListPassengerQueryByIdsResp{}, nil
}

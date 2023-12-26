package passenger

import (
	"context"
	"go-zero-12306/app/user-service/cmd/rpc/user"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SavePassengerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSavePassengerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SavePassengerLogic {
	return &SavePassengerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SavePassengerLogic) SavePassenger(req *types.SavePassengerReq) (*types.SavePassengerResp, error) {
	_, err := l.svcCtx.UserRpc.SavePassenger(l.ctx, &user.SavePassengerReq{
		Id:           req.Id,
		RealName:     req.RealName,
		IdType:       req.IdType,
		IdCard:       req.IdCard,
		DiscountType: req.DiscountType,
		Phone:        req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &types.SavePassengerResp{}, nil
}

package passenger

import (
	"context"
	"go-zero-12306/app/user-service/cmd/rpc/user"

	"go-zero-12306/app/user-service/cmd/api/internal/svc"
	"go-zero-12306/app/user-service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePassengerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePassengerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePassengerLogic {
	return &UpdatePassengerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePassengerLogic) UpdatePassenger(req *types.UpdatePassengerReq) (*types.UpdatePassengerResp, error) {
	_, err := l.svcCtx.UserRpc.UpdatePassenger(l.ctx, &user.SavePassengerReq{
		Id:           req.Id,
		RealName:     req.RealName,
		IdType:       req.IdType,
		IdCard:       req.IdCard,
		DiscountType: req.DiscountType,
		Phone:        req.Phone,
	})
	if err != nil {
		return &types.UpdatePassengerResp{}, err
	}
	return &types.UpdatePassengerResp{}, nil
}

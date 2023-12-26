package logic

import (
	"context"
	"github.com/jinzhu/copier"

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
	tPassenger0s, err := l.svcCtx.Passenger0Model.GetPassengerByUserName(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	set := ConvertIntSlice2Map(in.Ids)
	for i := 0; i < len(tPassenger0s); i++ {
		if _, ok := set[tPassenger0s[i].Id]; !ok {
			tPassenger0s = append(tPassenger0s[:i], tPassenger0s[i+1:]...)
			i--
		}
	}
	var list []*pb.ListPassengerQueryByIds
	_ = copier.Copy(&list, &tPassenger0s)
	return &pb.ListPassengerQueryByIdsResp{
		List: list,
	}, nil
}

func ConvertIntSlice2Map(s1 []int64) map[int64]bool {
	set := make(map[int64]bool, len(s1))
	for _, v := range s1 {
		set[v] = true
	}
	return set
}

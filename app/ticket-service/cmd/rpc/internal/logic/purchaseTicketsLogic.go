package logic

import (
	"context"
	"fmt"
	godisson "github.com/cheerego/go-redisson"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"
	"go-zero-12306/app/ticket-service/model/tTrain"
	"go-zero-12306/common/constant"
	"strconv"
)

type PurchaseTicketsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPurchaseTicketsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PurchaseTicketsLogic {
	return &PurchaseTicketsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
	购买车票
*/

func (l *PurchaseTicketsLogic) PurchaseTickets(in *pb.PurchaseTicketsReq) (*pb.PurchaseTicketsResp, error) {
	redisClient := l.svcCtx.RedisClient
	// 1. 责任链模式，参数验证
	// 2. 获取RLock锁
	lockKey := fmt.Sprintf(constant.LOCK_PURCHASE_TICKETS, in.TrainId)
	lock := godisson.NewGodisson(redisClient).NewRLock(lockKey)
	lock.Lock()
	defer lock.Unlock()
	// 3. 购票
	return l.executePurchaseTickets(in)
}

// 实际购票
func (l *PurchaseTicketsLogic) executePurchaseTickets(in *pb.PurchaseTicketsReq) (*pb.PurchaseTicketsResp, error) {
	// 开启事务
	if err := l.svcCtx.TTicketModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		cacheLogic := NewDistributedCacheLogic(l.ctx, l.svcCtx)
		key := constant.TRAIN_INFO + in.TrainId
		var trainDO = tTrain.TTrain{}
		// 从缓存中获取
		err := cacheLogic.SafeGet(key, &trainDO, nil, func() (interface{}, error) {
			builder := l.svcCtx.TTrainModel.SelectBuilder()
			parseInt, _ := strconv.ParseInt(in.TrainId, 10, 64)
			trainDO, err := l.svcCtx.TTrainModel.SelectById(l.ctx, builder, parseInt)
			if errors.Is(err, sqlx.ErrNotFound) {
				return nil, nil
			}
			if err != nil {
				return nil, err
			}
			return trainDO, nil
		}, nil, constant.ADVANCE_TICKET_DAY)
		if err != nil {
			return err
		}
		// todo:
		return nil
	}); err != nil {

	}
}

package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"go-zero-12306/app/ticket-service/cmd/rpc/pb"
	userPb "go-zero-12306/app/user-service/cmd/rpc/pb"
	"strconv"
	"sync"
)

type TrainSeatTypeSelector struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTrainSeatTypeSelector(ctx context.Context, svcCtx *svc.ServiceContext) *TrainSeatTypeSelector {
	return &TrainSeatTypeSelector{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TrainSeatTypeSelector) Select(trainType int64, requestParam *pb.PurchaseTicketsReq) ([]*pb.TrainPurchaseTicketRespDTO, error) {
	passengerDetails := requestParam.Passengers
	seatTypeMap := make(map[int64][]*pb.PurchaseTicketPassengerDetailDTO)
	for _, detail := range passengerDetails {
		seatTypeMap[detail.SeatType] = append(seatTypeMap[detail.SeatType], detail)
	}
	actualResult := make([]*pb.TrainPurchaseTicketRespDTO, 0)
	// 锁
	var wg sync.WaitGroup
	seatsChan := make(chan []*pb.TrainPurchaseTicketRespDTO, len(seatTypeMap))

	//  座位购买
	if len(seatTypeMap) > 0 {
		for seatType, passengerDetails := range seatTypeMap {
			wg.Add(1)
			go func(st int64, psd []*pb.PurchaseTicketPassengerDetailDTO) {
				defer wg.Done()
				result, err := l.distributeSeats(trainType, seatType, requestParam, passengerDetails)
				if err != nil {
					return
				}
				// 向chan发送
				seatsChan <- result
			}(seatType, passengerDetails)
		}
	} else {
		for seatType, passengerSeatDetails := range seatTypeMap {
			result, _ := l.distributeSeats(trainType, seatType, requestParam, passengerSeatDetails)
			actualResult = append(actualResult, result...)
		}
	}

	wg.Wait()
	close(seatsChan)
	for seats := range seatsChan {
		actualResult = append(actualResult, seats...)
	}
	if len(actualResult) == 0 || len(actualResult) != len(passengerDetails) {
		return nil, errors.New("站点余票不足，请尝试更换座位类型或选择其它站点")
	}
	// 获取passengers的ID
	passengerIds := make([]int64, 0)
	for _, resp := range actualResult {
		ids, _ := strconv.ParseInt(resp.PassengerId, 10, 64)
		passengerIds = append(passengerIds, ids)
	}
	//TODO:
	passengerRemoteResult, err := l.svcCtx.UserRpc.ListPassengerQueryByIds(l.ctx, &userPb.ListPassengerQueryByIdsReq{
		Username: requestParam.Username,
		Ids:      passengerIds,
	})
	if err != nil {
		return nil, errors.New("用户服务远程调用查询乘车人相信信息错误")
	}
	for i := range actualResult {
		for _, passenger := range passengerRemoteResult.List {
			if actualResult[i].PassengerId == passenger.Id {
				actualResult[i].IdCard = passenger.IdCard
				actualResult[i].Phone = passenger.Phone
				actualResult[i].UserType = passenger.DiscountType
				actualResult[i].IdType = passenger.IdType
				actualResult[i].RealName = passenger.RealName
				break
			}
		}
		// 查询票价
		builder := l.svcCtx.TTrainStationPriceModel.SelectBuilder()
		builder = builder.Where("train_id = ? and departure = ? and arrival = ? and seat_type = ?", requestParam.TrainId, requestParam.Departure, requestParam.Arrival, actualResult[i].SeatType)
		trainStationPriceDO, err := l.svcCtx.TTrainStationPriceModel.SelectOne(l.ctx, builder)
		if err != nil {
			return nil, err
		}
		actualResult[i].Amount = trainStationPriceDO.Price.Int64
	}
	// 锁定座位
	l.lockSeat(requestParam.TrainId, requestParam.Departure, requestParam.Arrival, actualResult)

	return actualResult, nil
}

func (l *TrainSeatTypeSelector) distributeSeats(trainType, seatType int64,
	requestParam *pb.PurchaseTicketsReq,
	passengerSeatDetails []*pb.PurchaseTicketPassengerDetailDTO) ([]*pb.TrainPurchaseTicketRespDTO, error) {
	// TODO: 座位选择 ，logic
	return nil, nil
}

/**
 * 锁定选中以及沿途车票状态
 *
 * @param trainId                     列车 ID
 * @param departure                   出发站
 * @param arrival                     到达站
 * @param trainPurchaseTicketRespList 乘车人以及座位信息
 */
func (l *TrainSeatTypeSelector) lockSeat(trainId, departure, arrival string, trainPurchaseTicketRespList []*pb.TrainPurchaseTicketRespDTO) {
	//TODO: 锁定座位,logic
}

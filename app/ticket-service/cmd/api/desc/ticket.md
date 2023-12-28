### 1. "查询车站&amp;城市站点集合信息"

1. route definition

- Url: /api/ticket-service/region-station/query
- Method: GET
- Request: `RegionStationQueryReqDTO`
- Response: `RegionStationQueryResp`

2. request definition



```golang
type RegionStationQueryReqDTO struct {
	QuerryType int64 `form:"querryType"` // 查询方式
	Name string `form:"name"` //名称
}
```


3. response definition



```golang
type RegionStationQueryResp struct {
	List []RegionStationQueryRespDTO `json:"list"`
}
```

### 2. "查询车站站点集合信息"

1. route definition

- Url: /api/ticket-service/station/all
- Method: GET
- Request: `StationQueryReq`
- Response: `StationQueryRespDTO`

2. request definition



```golang
type StationQueryReq struct {
}
```


3. response definition



```golang
type StationQueryRespDTO struct {
	Name string `json:"name"` //名称
	Code string `json:"code"` // 地区编码
	Spell string `json:"spell"` // 拼音
	RegionName string `json:"regionName"` // 城市名称
}
```

### 3. "座位重置"

1. route definition

- Url: /api/ticket-service/temp/seat/reset
- Method: POST
- Request: `PurchaseTicketsReq`
- Response: `PurchaseTicketsResp`

2. request definition



```golang
type PurchaseTicketsReq struct {
	TrainId string `form:"trainId"`
}
```


3. response definition



```golang
type PurchaseTicketsResp struct {
}
```

### 4. "根据条件查询车票"

1. route definition

- Url: /api/ticket-service/ticket/query
- Method: GET
- Request: `TicketPageQueryReqDTO`
- Response: `TicketPageQueryRespDTO`

2. request definition



```golang
type TicketPageQueryReqDTO struct {
	FromStation string `form:"fromStation"` //出发地 Code
	ToStation string `form:"toStation"` //目的地 Code
	DepartureDate string `form:"departureDate"` //出发日期
	Departure string `form:"departure"` //出发站点
	Arrival string `form:"arrival"` //到达站点
}
```


3. response definition



```golang
type TicketPageQueryRespDTO struct {
	TrainList []TicketListDTO `json:"trainList"` //车次集合数据
	TrainBrandList []int64 `json:"trainBrandList"` // 车次类型：D-动车 Z-直达 复兴号等
	DepartureStationList []string `json:"departureStationList"` //出发车站
	ArrivalStationList []string `json:"arrivalStationList"` // 到达车站
	SeatClassTypeList []int64 `json:"seatClassTypeList"` //车次席别
}
```

### 5. "购买车票"

1. route definition

- Url: /api/ticket-service/ticket/purchase
- Method: POST
- Request: `PurchaseTicketReqDTO`
- Response: `TicketPurchaseRespDTO`

2. request definition



```golang
type PurchaseTicketReqDTO struct {
	TrainId string `json:"trainId"` //车次 ID
	Passengers []PurchaseTicketPassengerDetailDTO `json:"passengers"` //乘车人
	ChooseSeats []string `json:"chooseSeats"` //选择座位
	Departure string `json:"departure"` //出发站点
	Arrival string `json:"arrival"` //到达站点
}
```


3. response definition



```golang
type TicketPurchaseRespDTO struct {
	OrderSn string `json:"orderSn"` //订单号
	TicketOrderDetails []TicketOrderDetailRespDTO `json:"ticketOrderDetails"` //乘车人订单详情
}
```

### 6. "购买车票v2"

1. route definition

- Url: /api/ticket-service/ticket/purchase/v2
- Method: POST
- Request: `PurchaseTicketReqDTO`
- Response: `TicketPurchaseRespDTO`

2. request definition



```golang
type PurchaseTicketReqDTO struct {
	TrainId string `json:"trainId"` //车次 ID
	Passengers []PurchaseTicketPassengerDetailDTO `json:"passengers"` //乘车人
	ChooseSeats []string `json:"chooseSeats"` //选择座位
	Departure string `json:"departure"` //出发站点
	Arrival string `json:"arrival"` //到达站点
}
```


3. response definition



```golang
type TicketPurchaseRespDTO struct {
	OrderSn string `json:"orderSn"` //订单号
	TicketOrderDetails []TicketOrderDetailRespDTO `json:"ticketOrderDetails"` //乘车人订单详情
}
```

### 7. "取消车票订单"

1. route definition

- Url: /api/ticket-service/ticket/cancel
- Method: POST
- Request: `CancelTicketOrderReqDTO`
- Response: `CancelTicketOrderResp`

2. request definition



```golang
type CancelTicketOrderReqDTO struct {
	OrderSn string `json:"orderSn"` // 订单号
}
```


3. response definition



```golang
type CancelTicketOrderResp struct {
}
```

### 8. "支付单详情查询"

1. route definition

- Url: /api/ticket-service/ticket/pay/query
- Method: GET
- Request: `GetPayInfoReq`
- Response: `PayInfoRespDTO`

2. request definition



```golang
type GetPayInfoReq struct {
	OrderSn string `form:"orderSn"`
}
```


3. response definition



```golang
type PayInfoRespDTO struct {
	OrderSn string `json:"orderSn"` // 订单号
	TotalAmount int64 `json:"totalAmount"` //支付总金额
	Status int64 `json:"status"` //支付状态
	GmtPayment string `json:"gmtPayment"` //支付时间
}
```

### 9. "公共退款接口"

1. route definition

- Url: /api/ticket-service/ticket/refund
- Method: POST
- Request: `RefundTicketReqDTO`
- Response: `RefundTicketRespDTO`

2. request definition



```golang
type RefundTicketReqDTO struct {
	OrderSn string `json:"orderSn"` // 订单号
	Type int64 `json:"type"` // 退款类型 0 部分退款 1 全部退款
	SubOrderRecordIdReqList []string `json:"subOrderRecordIdReqList"` // 部分退款子订单记录id集合
}
```


3. response definition



```golang
type RefundTicketRespDTO struct {
}
```

### 10. "根据列车 ID 查询站点信息"

1. route definition

- Url: /api/ticket-service/train-station/query
- Method: GET
- Request: `ListTrainStationQueryReq`
- Response: `TrainStationQueryResp`

2. request definition



```golang
type ListTrainStationQueryReq struct {
	TrainId string `form:"trainId"`
}
```


3. response definition



```golang
type TrainStationQueryResp struct {
	List []TrainStationQueryRespDTO `json:"list"`
}
```


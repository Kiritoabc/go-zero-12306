### 1. "根据订单号查询车票订单"

1. route definition

- Url: /api/order-service/order/ticket/query
- Method: GET
- Request: `QueryTicketOrderByOrderSnReq`
- Response: `QueryTicketOrderByOrderSnResp`

2. request definition



```golang
type QueryTicketOrderByOrderSnReq struct {
	OrderSn string `json:"orderSn"`
}
```


3. response definition



```golang
type QueryTicketOrderByOrderSnResp struct {
	OrderSn string `json:"orderSn"` // 订单号
	TrainId int64 `json:"trainId"` // 列车ID
	Departure string `json:"departure"` // 出发站点
	Arrival string `json:"arrival"` // 到达站点
	RidingDate string `json:"ridingDate"` // 乘车日期
	OrderTime string `json:"orderTime"` // 订票日期
	TrainNumber string `json:"trainNumber"` // 列车车次
	DepartureTime string `json:"departureTime"` // 出发时间
	ArrivalTime string `json:"arrivalTime"` // 到达时间
	PassengerDetails []TicketOrderPassengerDetailRespDTO `json:"passengerDetails"` //乘车人订单信息
}
```

### 2. "根据子订单记录id查询车票子订单详情"

1. route definition

- Url: /api/order-service/order/item/ticket/query
- Method: GET
- Request: `QueryTicketItemOrderByIdReq`
- Response: `QueryTicketItemOrderByIdResp`

2. request definition



```golang
type QueryTicketItemOrderByIdReq struct {
	OrderSn string `json:"orderSn"` // 订单号
	OrderItemRecordIds []int64 `json:"orderItemRecordIds"` // 子订单记录id
}
```


3. response definition



```golang
type QueryTicketItemOrderByIdResp struct {
	List []TicketOrderPassengerDetailRespDTO `json:"list"`
}
```

### 3. "分页查询车票订单"

1. route definition

- Url: /api/order-service/order/ticket/page
- Method: GET
- Request: `PageTicketOrderReq`
- Response: `PageTicketOrderResp`

2. request definition



```golang
type PageTicketOrderReq struct {
	UserId string `json:"userId"` // 用户唯一标识
	StatusType int64 `json:"statusType"` // 状态类型 0：未完成 1：未出行 2：历史订单
}
```


3. response definition



```golang
type PageTicketOrderResp struct {
	Current int64 `json:"current"` // 当前页
	Size int64 `json:"size"` // 每页显示条数
	Total int64 `json:"total"` // 总数
	Records []QueryTicketItemOrderByIdResp `json:"records"` //查询数据列表
}
```

### 4. "分页查询本人车票订单"

1. route definition

- Url: /api/order-service/order/ticket/self/page
- Method: GET
- Request: `TicketOrderSelfPageQueryReq`
- Response: `TicketOrderSelfPageQueryResp`

2. request definition



```golang
type TicketOrderSelfPageQueryReq struct {
	Current int64 `json:"current"` // 当前页
	Size int64 `json:"size"` // 每页显示的条数
}
```


3. response definition



```golang
type TicketOrderSelfPageQueryResp struct {
	Current int64 `json:"current"` // 当前页
	Size int64 `json:"size"` // 每页显示条数
	Total int64 `json:"total"` // 总数
	Records []TicketOrderDetailSelfRespDTO `json:"records"` //查询数据列表
}
```

### 5. "车票订单创建"

1. route definition

- Url: /api/order-service/order/ticket/create
- Method: POST
- Request: `CreateTicketOrderReq`
- Response: `CreateTicketOrderResp`

2. request definition



```golang
type CreateTicketOrderReq struct {
	UserId int64 `json:"userId"` // 用户Id
	UserName string `json:"userName"` // 用户名
	TrainId int64 `json:"trainId"` // 车次ID
	Departure string `json:"departure"` // 出发站点
	Arrival string `json:"arrival"` // 到达站点
	Source int64 `json:"source"` // 订单来源
	OrderTime string `json:"orderTime"` // 下单时间
	RidingDate string `json:"ridingDate"` // 乘车日期
	TrainNumber string `json:"trainNumber"` // 乘车车次
	DepartureTime string `json:"departureTime"` // 出发时间
	ArrivalTime string `json:"arrivalTime"` // 到达时间
	TicketOrderItems []TicketOrderItemCreateReqDTO `json:"ticketOrderItems"` // 订单明细
}
```


3. response definition



```golang
type CreateTicketOrderResp struct {
	Data string `json:"data"` // 返回数据
}
```

### 6. "车票订单关闭"

1. route definition

- Url: /api/order-service/order/ticket/close
- Method: POST
- Request: `CancelTicketOrderReq`
- Response: `CancelTicketOrderResp`

2. request definition



```golang
type CancelTicketOrderReq struct {
	OrderSn string `json:"orderSn"` // 订单号
}
```


3. response definition



```golang
type CancelTicketOrderResp struct {
	Bool bool `json:"bool"` //  是否成功
}
```

### 7. "车票订单取消"

1. route definition

- Url: /api/order-service/order/ticket/cancel
- Method: POST
- Request: `CancelTicketOrderReq`
- Response: `CancelTicketOrderResp`

2. request definition



```golang
type CancelTicketOrderReq struct {
	OrderSn string `json:"orderSn"` // 订单号
}
```


3. response definition



```golang
type CancelTicketOrderResp struct {
	Bool bool `json:"bool"` //  是否成功
}
```


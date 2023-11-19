// Code generated by goctl. DO NOT EDIT.
package types

type QueryTicketOrderByOrderSnReq struct {
	OrderSn string `json:"orderSn"`
}

type QueryTicketOrderByOrderSnResp struct {
	OrderSn          string                              `json:"orderSn"`          // 订单号
	TrainId          int64                               `json:"trainId"`          // 列车ID
	Departure        string                              `json:"departure"`        // 出发站点
	Arrival          string                              `json:"arrival"`          // 到达站点
	RidingDate       string                              `json:"ridingDate"`       // 乘车日期
	OrderTime        string                              `json:"orderTime"`        // 订票日期
	TrainNumber      string                              `json:"trainNumber"`      // 列车车次
	DepartureTime    string                              `json:"departureTime"`    // 出发时间
	ArrivalTime      string                              `json:"arrivalTime"`      // 到达时间
	PassengerDetails []TicketOrderPassengerDetailRespDTO `json:"passengerDetails"` //乘车人订单信息
}

type QueryTicketItemOrderByIdReq struct {
	OrderSn            string  `json:"orderSn"`            // 订单号
	OrderItemRecordIds []int64 `json:"orderItemRecordIds"` // 子订单记录id
}

type QueryTicketItemOrderByIdResp struct {
	List []TicketOrderPassengerDetailRespDTO `json:"list"`
}

type TicketOrderPassengerDetailRespDTO struct {
	Id             string `json:"id"`             // ID
	UserId         string `json:"userId"`         // 用户Id
	Username       string `json:"username"`       // 用户名
	SeatType       int64  `json:"seatType"`       // 席别类型
	CarriageNumber string `json:"carriageNumber"` // 车厢号
	SeatNumber     string `json:"seatNumber"`     // 座位号
	RealName       string `json:"realName"`       // 真实姓名
	IdType         int64  `json:"idType"`         // 证件类型
	IdCard         string `json:"idCard"`         // 证件号
	TicketType     int64  `json:"ticketType"`     // 车票类型 0：成人 1：儿童 2：学生 3：残疾军人
	Amount         int64  `json:"amount"`         // 订单金额
	Status         int64  `json:"status"`         // 车票状态
	StatusName     string `json:"statusName"`     // 车票状态名称
}

type PageTicketOrderReq struct {
	UserId     string `json:"userId"`     // 用户唯一标识
	StatusType int64  `json:"statusType"` // 状态类型 0：未完成 1：未出行 2：历史订单
}

type PageTicketOrderResp struct {
	Current int64                          `json:"current"` // 当前页
	Size    int64                          `json:"size"`    // 每页显示条数
	Total   int64                          `json:"total"`   // 总数
	Records []QueryTicketItemOrderByIdResp `json:"records"` //查询数据列表
}

type TicketOrderSelfPageQueryReq struct {
	Current int64 `json:"current"` // 当前页
	Size    int64 `json:"size"`    // 每页显示的条数
}

type TicketOrderSelfPageQueryResp struct {
	Current int64                          `json:"current"` // 当前页
	Size    int64                          `json:"size"`    // 每页显示条数
	Total   int64                          `json:"total"`   // 总数
	Records []TicketOrderDetailSelfRespDTO `json:"records"` //查询数据列表
}

type TicketOrderDetailSelfRespDTO struct {
	Departure      string `json:"departure"`      // 出发站点
	Arrival        string `json:"arrival"`        // 到达站点
	RidingDate     string `json:"ridingDate"`     // 乘车日期
	TrainNumber    string `json:"trainNumber"`    // 乘车车次
	DepartureTime  string `json:"departureTime"`  // 出发时间
	ArrivalTime    string `json:"arrivalTime"`    // 到达时间
	SeatType       int64  `json:"seatType"`       // 席别类型
	CarriageNumber string `json:"carriageNumber"` // 车厢号
	SeatNumber     string `json:"seatNumber"`     // 座位号
	RealName       string `json:"realName"`       // 真实姓名
	TicketType     int64  `json:"ticketType"`     // 车票类型 0：成人 1：儿童 2：学生 3：残疾军人
	Amount         int64  `json:"amount"`         // 订单金额
}

type CreateTicketOrderReq struct {
	UserId           int64                         `json:"userId"`           // 用户Id
	UserName         string                        `json:"userName"`         // 用户名
	TrainId          int64                         `json:"trainId"`          // 车次ID
	Departure        string                        `json:"departure"`        // 出发站点
	Arrival          string                        `json:"arrival"`          // 到达站点
	Source           int64                         `json:"source"`           // 订单来源
	OrderTime        string                        `json:"orderTime"`        // 下单时间
	RidingDate       string                        `json:"ridingDate"`       // 乘车日期
	TrainNumber      string                        `json:"trainNumber"`      // 乘车车次
	DepartureTime    string                        `json:"departureTime"`    // 出发时间
	ArrivalTime      string                        `json:"arrivalTime"`      // 到达时间
	TicketOrderItems []TicketOrderItemCreateReqDTO `json:"ticketOrderItems"` // 订单明细
}

type TicketOrderItemCreateReqDTO struct {
	CarriageNumber string `json:"carriageNumber"` // 车厢号
	SeatType       int64  `json:"seatType"`       // 席别类型
	SeatNumber     string `json:"seatNumber"`     // 座位号
	PassengerId    string `json:"passengerId"`    // 乘车人ID
	RealName       string `json:"realName"`       // 真实姓名
	IdType         int64  `json:"idType"`         // 证件类型
	IdCard         string `json:"idCard"`         // 证件号
	Phone          string `json:"phone"`          // 手机号
	Amount         string `json:"amount"`         // 订单金额
	TicketType     int64  `json:"ticketType"`     // 车票类型
}

type CreateTicketOrderResp struct {
	Data string `json:"data"` // 返回数据
}

type CancelTicketOrderReq struct {
	OrderSn string `json:"orderSn"` // 订单号
}

type CancelTicketOrderResp struct {
	Bool bool `json:"bool"` //  是否成功
}

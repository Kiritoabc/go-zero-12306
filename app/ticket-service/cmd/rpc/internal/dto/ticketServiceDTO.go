package dto

type TicketListDTO struct {
	// 列车 ID
	TrainId string `json:"trainId"`

	// 车次
	TrainNumber string `json:"trainNumber"`

	// 触发时间
	DepartureTime string `json:"arrivalTime"`

	// 到达时间
	ArrivalTime string `json:"arrivalTime"`

	// 历时
	Duration string `json:"duration"`

	// 到达天数
	DaysArrived int64 `json:"daysArrived"`

	// 出发站点
	Departure string `json:"departure"`

	// 到达站点
	Arrival string `json:"arrival"`

	// 始发站标识
	DepartureFlag bool `json:"departureFlag"`

	// 终点站标识
	ArrivalFlag bool `json:"arrivalFlag"`

	//列车类型 0：高铁 1：动车 2：普通车
	TrainType int `json:"trainType"`

	// 可售时间
	SaleTime string `json:"saleTime"`

	// 销售状态 0：可售 1：不可售 2：未知
	SaleStatus int64 `json:"saleStatus"`

	// 列车品牌类型 0：GC-高铁/城际 1：D-动车 2：Z-直达 3：T-特快 4：K-快速 5：其他 6：复兴号 7：智能动车组
	TrainBrand string `json:"trainBrand"`

	/**
	 * 列车标签集合 0：复兴号 1：智能动车组 2：静音车厢 3：支持选铺
	 */
	TrainTags []string `json:"trainTags"`

	// 写别实体集合
	SeatClassList []SeatClassDTO `json:"seatClassList"`
}

type SeatClassDTO struct {
	// 席别类型
	Type int `json:"type"`

	// 席别数量
	Quantity int `json:"quantity"`

	// 席别价格
	Price float64 `json:"price"` // TODO: 注意Go标准库没有内置BigDecimal类型，需要引入第三方库如"github.com/shopspring/decimal"

	// 席别候补标识
	Candidate bool `json:"candidate"`
}

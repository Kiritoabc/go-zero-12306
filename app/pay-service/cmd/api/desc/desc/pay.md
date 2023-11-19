### 1. "公共支付接口,对接常用支付方式，比如：支付宝、微信以及银行卡等"

1. route definition

- Url: /api/pay-service/pay/create
- Method: POST
- Request: `PayReq`
- Response: `PayResp`

2. request definition



```golang
type PayReq struct {
	OutOrderSn string `json:"outOrderSn"` // 子订单号
	TotalAmount float64 `json:"totalAmount"` //订单总金额,单位为元，精确到小数点后两位，取值范围：[0.01,100000000]
	Subject string `json:"subject"` // 订单标题 ，注意：不可使用特殊字符，如 /，=，&amp; 等
}
```


3. response definition



```golang
type PayResp struct {
	Body string `json:"body"` // 调用支付返回信息
}
```

### 2. "跟据订单号查询支付单详情"

1. route definition

- Url: /api/pay-service/pay/query/order-sn
- Method: GET
- Request: `GetPayInfoByOrderSnReq`
- Response: `GetPayInfoByOrderSnResp`

2. request definition



```golang
type GetPayInfoByOrderSnReq struct {
	OrderSn string `json:"orderSn"` //订单号
}
```


3. response definition



```golang
type GetPayInfoByOrderSnResp struct {
	OrderSn string `json:"orderSn"` //订单号
	TotalAmount int64 `json:"totalAmount"` // 支付总金额
	Status int64 `json:"status"` // 支付状态
	GmtPayment string `json:"gmtPayment"` // 支付时间
}
```

### 3. "跟据支付流水号查询支付单详情"

1. route definition

- Url: /api/pay-service/pay/query/pay-sn
- Method: GET
- Request: `GetPayInfoByPaySnReq`
- Response: `GetPayInfoByPaySnResp`

2. request definition



```golang
type GetPayInfoByPaySnReq struct {
	PaySn string `json:"paySn"` // 支付流水号
}
```


3. response definition



```golang
type GetPayInfoByPaySnResp struct {
	OrderSn string `json:"orderSn"` //订单号
	TotalAmount int64 `json:"totalAmount"` // 支付总金额
	Status int64 `json:"status"` // 支付状态
	GmtPayment string `json:"gmtPayment"` // 支付时间
}
```

### 4. "公共退款接口,后续为了方便开发系列退款相关接口，已迁移 {@link RefundController#commonRefund(RefundReqDTO)}"

1. route definition

- Url: /api/pay-service/refund
- Method: POST
- Request: `RefundReq`
- Response: `RefundResp`

2. request definition



```golang
type RefundReq struct {
	OrderSn string `json:"orderSn"` // 订单号
	RefundTypeEnum string `json:"refundTypeEnum"` // 退款类型枚举 （生成代码后修改）
	RefundAmount int64 `json:"refundAmount"` // 退款金额
	RefundDetailReqDTOList []TicketOrderPassengerDetailRespDTO `json:"refundDetailReqDtoList"`
}
```


3. response definition



```golang
type RefundResp struct {
}
```

### 5. "支付宝回调,调用支付宝支付后，支付宝会调用此接口发送支付结果"

1. route definition

- Url: /api/pay-service/callback/alipay
- Method: POST
- Request: `CallbackAlipayReq`
- Response: `CallbackAlipayResp`

2. request definition



```golang
type CallbackAlipayReq struct {
}
```


3. response definition



```golang
type CallbackAlipayResp struct {
}
```

### 6. "公共退款接口"

1. route definition

- Url: /api/pay-service/common/path
- Method: POST
- Request: `RefundReq`
- Response: `RefundResp`

2. request definition



```golang
type RefundReq struct {
	OrderSn string `json:"orderSn"` // 订单号
	RefundTypeEnum string `json:"refundTypeEnum"` // 退款类型枚举 （生成代码后修改）
	RefundAmount int64 `json:"refundAmount"` // 退款金额
	RefundDetailReqDTOList []TicketOrderPassengerDetailRespDTO `json:"refundDetailReqDtoList"`
}
```


3. response definition



```golang
type RefundResp struct {
}
```


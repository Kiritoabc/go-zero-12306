package result

type ResponseSuccessBean struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{"200", "OK", data}
}

type ResponseErrorBean struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func Error(errCode string, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}

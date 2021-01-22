package server_context

// Response http响应对象
type Response struct {
	Code    int32       `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// NewError 创建一个错误
func NewError(code int32, data interface{}) *Response {
	key, ok := errsMap[code]
	if ok == false {
		code = -1
		key = errsMap[-1]
	}
	// 构建返回消息
	return &Response{
		Code:    code,
		Message: key,
		Data:    data,
	}
}

// NewSuccess 创建一个成功消息
func NewSuccess(data interface{}) *Response {
	return NewError(1, data)
}

// GetServiceError 用于服务返回错误使用
func GetServiceError(code int32) (int32, string) {
	msg := NewError(code, nil)
	return msg.Code, msg.Message
}

var errsMap = map[int32]string{
	-1: "Error",
	1:  "Success",
}


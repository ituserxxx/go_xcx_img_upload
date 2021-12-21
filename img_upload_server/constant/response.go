package constant

type ResponseCode int

const (
	SUCCESS         ResponseCode = 1   // 成功
	ParamsError     ResponseCode = 900 // 参数校验失败
	LogicError      ResponseCode = 700 // 处理逻辑出错
	MiddleAuthError ResponseCode = 400 // 中间件验证逻辑出错

)

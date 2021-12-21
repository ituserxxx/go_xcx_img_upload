package response

import (
	"github.com/gogf/gf/net/ghttp"
	"img-xcx-tool/constant"
)

type ps struct {
	Code constant.ResponseCode `json:"code"`
	Msg  string                `json:"msg"`
	Data interface{}           `json:"data"`
}

func Succ(r *ghttp.Request, data interface{}) {

	_ = r.Response.WriteJsonExit(ps{
		Code: constant.SUCCESS,
		Msg:  "ok",
		Data: data,
	})
}

func Err(r *ghttp.Request, ErrCode constant.ResponseCode, message string) {

	_ = r.Response.WriteJsonExit(ps{
		Code: ErrCode,
		Msg:  message,
		Data: nil,
	})
}

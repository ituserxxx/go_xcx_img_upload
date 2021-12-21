package api

import (
	"github.com/gogf/gf/net/ghttp"
	"img-xcx-tool/app/in_out"
	"img-xcx-tool/app/service"
	"img-xcx-tool/constant"
	"img-xcx-tool/library/response"
)

var Login *loginApi
type loginApi struct {}

func (l *loginApi)Login(r *ghttp.Request){
	var req *in_out.LoginReq
	if err := r.Parse(&req); err != nil {
		response.Err(r, constant.ParamsError, err.Error())
	}
	resp ,err := service.Login.Login(req)
	if err != nil {
		response.Err(r, constant.LogicError, err.Error())
	}
	response.Succ(r,resp)
}

package api

import (
	"github.com/gogf/gf/net/ghttp"
	"img-xcx-tool/app/in_out"
	"img-xcx-tool/app/service"
	"img-xcx-tool/constant"
	"img-xcx-tool/library/response"
)
var UserImage *userImageApi
type userImageApi struct{}

func (ui *userImageApi) Add(r *ghttp.Request) {
	var req *in_out.AddUserImage
	if err := r.Parse(&req);err !=nil{
		response.Err(r,constant.LogicError,err.Error())
	}
	err := service.User.Add(req)
	if err != nil {
		response.Err(r,constant.LogicError,err.Error())
	}
	response.Succ(r,"")
}

func (ui *userImageApi) List(r *ghttp.Request) {
	var req *in_out.UserImageListReq
	if err := r.Parse(&req);err !=nil{
		response.Err(r,constant.LogicError,err.Error())
	}
	l,err := service.User.UserImageList(req)
	if err != nil {
		response.Err(r,constant.LogicError,err.Error())
	}
	response.Succ(r,l)
}

package api

import (
	"github.com/gogf/gf/net/ghttp"
	"img-xcx-tool/app/in_out"
	"img-xcx-tool/app/service"
	"img-xcx-tool/constant"
	"img-xcx-tool/library/response"
)

var ImageCategory *imageCategoryApi

type imageCategoryApi struct{}

func (ic *imageCategoryApi) Add(r *ghttp.Request) {
	var req *in_out.AddCategoryReq
	if err := r.Parse(&req);err!=nil{
		response.Err(r,constant.ParamsError,err.Error())
	}
	err := service.ImageCategory.Add(req)
	if err != nil {
		response.Err(r,constant.LogicError,err.Error())
	}
	response.Succ(r,"")
}

func (ic *imageCategoryApi) Update(r *ghttp.Request) {
	var req *in_out.UpdateCategoryReq
	if err := r.Parse(&req);err!=nil{
		response.Err(r,constant.ParamsError,err.Error())
	}
	err := service.ImageCategory.Update(req)
	if err != nil {
		response.Err(r,constant.LogicError,err.Error())
	}
	response.Succ(r,"")
}

package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"img-xcx-tool/app/in_out"
	"img-xcx-tool/library/response"
	"img-xcx-tool/library/tools"
)

var QiniuApi *qiniuApi

type qiniuApi struct {
}
func (qa *qiniuApi) GetToken(r *ghttp.Request) {
	var resp = &in_out.QiniuResponse{}
	resp.Token = tools.GetUploadToken()
	resp.Domain = g.Config().GetString("qiniu.domain")
	//resp.Region = g.Config().GetString("qiniu.region")
	resp.Region = "https://up-z0.qiniup.com"
	response.Succ(r, resp)
}

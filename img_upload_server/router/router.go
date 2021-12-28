package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"img-xcx-tool/app/api"
	"img-xcx-tool/library/wechat"
)

func init() {
	s := g.Server()
	s.BindHandler("/xxx", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
		gzh := wechat.GetGzh()
		az, err := gzh.GetAccessToken()
		if err != nil {
			r.Response.Write(err.Error())
		} else {
			r.Response.Write(az)
		}
	})
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.ALL("/login", api.Login.Login)
		group.ALL("/user/image/list", api.UserImage.UserImageList)
		group.ALL("/user/image/del", api.UserImage.DelImg)
		group.ALL("/user/image/love", api.UserImage.LoveImg)
		group.ALL("/user/image/add", api.UserImage.Add)
		group.ALL("/upload/token", api.QiniuApi.GetToken)

	})
}

package router

import (
	"img-xcx-tool/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.ALL("/login", api.Login.Login)
		group.ALL("/user/image/list", api.UserImage.UserImageList)
		group.ALL("/user/image/del", api.UserImage.DelImg)
		group.ALL("/user/image/love", api.UserImage.LoveImg)
		group.ALL("/user/image/add", api.UserImage.Add)
		group.ALL("/upload/token", api.QiniuApi.GetToken)

	})
}

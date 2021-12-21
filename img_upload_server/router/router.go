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
		group.ALL("/image_category/add", api.ImageCategory.Add)
		group.ALL("/image_category/update", api.ImageCategory.Update)
		group.ALL("/user/image/list", api.UserImage.List)
		group.ALL("/user/image/add", api.UserImage.Add)
		group.ALL("/upload/token", api.QiniuApi.GetToken)

	})
}

package service

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"img-xcx-tool/app/dao"
	"img-xcx-tool/app/in_out"
	"img-xcx-tool/app/model"
	"img-xcx-tool/library/wechat"
)

var Login *login
type login struct {}

func(l *login)Login(req *in_out.LoginReq)(int,error){
	encrInfo, err := wechat.Login(req.Code)
	if err != nil {
		return 0, gerror.New("Login 1" + err.Error())
	}
	u,err := dao.User.OmitEmpty().Where("openid =?",encrInfo.OpenID).FindOne()
	if err != nil {
		return 0, gerror.New("Login 2" + err.Error())
	}
	if u == nil{
		result,err := dao.User.OmitEmpty().Save(model.User{
			Openid:     encrInfo.OpenID,
			SessionKey: encrInfo.SessionKey,
			CreateTime: gtime.Now(),
		})
		if err != nil {
			return 0,err
		}
		id,err :=result.LastInsertId()
		if err != nil {
			return 0,err
		}
		return gconv.Int(id), nil
	}else{
		var uInfo model.User
		if err := u.Struct(&uInfo);err!=nil{
			return 0, err
		}
		return uInfo.Id, nil
	}
}

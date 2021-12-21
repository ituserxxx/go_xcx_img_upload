package wechat

import (
	"github.com/gogf/gf/frame/g"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
)

func GetMini() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	redisCache := &cache.RedisOpts{}
	cacheClient := cache.NewRedis(redisCache)
	wc.SetCache(cacheClient)
	cfg := &miniConfig.Config{
		AppID:     g.Config().GetString("mini_program.appid"),
		AppSecret: g.Config().GetString("mini_program.secret"),
		Cache:     nil,
	}
	return wc.GetMiniProgram(cfg)
}

// 登录
func Login(code string)(auth.ResCode2Session,error){
	var info auth.ResCode2Session
	mini := GetMini()
	isAuth := mini.GetAuth()
	info, err := isAuth.Code2Session(code)
	if err != nil || info.ErrCode != 0 {
		return info,err
	}
	return info,nil
}


// DecryptMsg 解密iv  data
func DecryptMsg(SessionKey, EncryptedData, Iv string) (*encryptor.PlainData, error) {
	mini := GetMini()
	encrypt := mini.GetEncryptor()
	info, err := encrypt.Decrypt(SessionKey, EncryptedData, Iv)
	if err != nil {
		return nil, err
	}
	return info, nil
}

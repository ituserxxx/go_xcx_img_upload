package service

import (
	"fmt"
	"img-xcx-tool/library/wechat"
	"testing"
)

func TestName(t *testing.T) {
	gzh := wechat.GetGzh()
	s, err := gzh.GetAccessToken()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s)
	}
}

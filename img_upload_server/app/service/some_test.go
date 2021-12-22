package service

import (
	"fmt"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	s1 :="https://qiniu.storage.heifengni.com/tmp/wFFNdeDGiR3i463c4d9db955123719b74304e4d9fd13.jpg"
	index := strings.LastIndex(s1, ".")
	s := []rune(s1)

	switch string(s[index:]) {
	case ".png",".jpg",".gif",".jpeg":
		fmt.Println(false)
	default:
		fmt.Println(true)
	}
}

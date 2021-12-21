package tools

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

func GetDateDiff(endTime , startTime string)int{
	s1 := gtime.New(startTime).Format("Y-m-d")
	s2 := gtime.New(endTime).Format("Y-m-d")
	ss2 := gtime.New(s2).Timestamp()
	ss1 := gtime.New(s1).Timestamp()
	return gconv.Int((ss2-ss1)/86400)

}

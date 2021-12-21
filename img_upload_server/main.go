package main

import (
	_ "img-xcx-tool/boot"
	_ "img-xcx-tool/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}

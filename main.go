package main

import (
	lib "joyconn-server-manage/lib"
	_ "joyconn-server-manage/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	lib.AddFuncMaps()
	web.Run()
}

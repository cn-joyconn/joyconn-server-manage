package main

import (
	"github.com/beego/beego/v2/server/web"
	lib "openvpn-server-manage/lib"
	_ "openvpn-server-manage/routers"
)

func main() {
	lib.AddFuncMaps()
	web.Run()
}

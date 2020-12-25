// Package routers defines application routes
// @APIVersion 1.0.0
// @Title OpenVPN API
// @Description REST API allows you to control and monitor your OpenVPN server
// @Contact adam.walach@gmail.com
// License Apache 2.0
// LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"strings"
	"unicode/utf8"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/config"

	controllers "joyconn-server-manage/controllers"
	baseController "joyconn-server-manage/controllers/base"
	accountController "joyconn-server-manage/controllers/account"
	appController "joyconn-server-manage/controllers/app"
	vpnController "joyconn-server-manage/controllers/vpn"
)

func init() {
	// val, err := config.String("ContextPath")
	// if (err!=nil){
	// 	val=""
	// }else if strings.HasSuffix(val, "/"){
	// 	val=val[:utf8.RuneCountInString(val) -1]
	// }
	web.SetStaticPath("/swagger", "swagger")
	web.Router("/", &controllers.MainController{})
	web.Router("/login", &accountController.LoginController{}, "get,post:Login")
	web.Router("/logout", &accountController.LoginController{}, "get:Logout")
	web.Router("/profile", &accountController.ProfileController{})
	web.Router("/settings", &appController.SettingsController{})
	web.Router("/ov/config", &vpnController.OVConfigController{})
	web.Router("/logs", &appController.LogsController{})

	web.Include(&vpnController.CertificatesController{})

	ns := web.NewNamespace("/api/v1",
			web.NSNamespace("/sysload",
				web.NSInclude(
						&baseController.APISysloadController{},
					),
			),
			web.NSNamespace("/signal",
				web.NSInclude(
						&baseController.APISignalController{},
					),
			),
		)
	web.AddNamespace(ns)
}

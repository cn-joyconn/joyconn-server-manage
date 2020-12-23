// Package routers defines application routes
// @APIVersion 1.0.0
// @Title OpenVPN API
// @Description REST API allows you to control and monitor your OpenVPN server
// @Contact adam.walach@gmail.com
// License Apache 2.0
// LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/beego/beego/v2/server/web"
	controllers "openvpn-server-manage/controllers"
)

func init() {
	web.SetStaticPath("/swagger", "swagger")
	web.Router("/", &controllers.MainController{})
	web.Router("/login", &controllers.LoginController{}, "get,post:Login")
	web.Router("/logout", &controllers.LoginController{}, "get:Logout")
	web.Router("/profile", &controllers.ProfileController{})
	web.Router("/settings", &controllers.SettingsController{})
	web.Router("/ov/config", &controllers.OVConfigController{})
	web.Router("/logs", &controllers.LogsController{})

	web.Include(&controllers.CertificatesController{})

	ns := web.NewNamespace("/api/v1",
			web.NSNamespace("/session",
				web.NSInclude(
							&controllers.APISessionController{},
					),
				),
			web.NSNamespace("/sysload",
				web.NSInclude(
						&controllers.APISysloadController{},
					),
			),
			web.NSNamespace("/signal",
				web.NSInclude(
						&controllers.APISignalController{},
					),
			),
		)
	web.AddNamespace(ns)
}

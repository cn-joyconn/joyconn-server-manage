package routers

import (
	"github.com/beego/beego/v2/server/web"
)

func init() {

	web.GlobalControllerRouter["openvpn-server-manage/controllers:APISessionController"] = append(
		web.GlobalControllerRouter["openvpn-server-manage/controllers:APISessionController"],
		web.ControllerComments{
				Method: "Get",
				Router: `/`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

				web.GlobalControllerRouter["openvpn-server-manage/controllers:APISessionController"] = append(web.GlobalControllerRouter["openvpn-server-manage/controllers:APISessionController"],
				web.ControllerComments{
			Method: "Kill",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

			web.GlobalControllerRouter["openvpn-server-manage/controllers:APISignalController"] = append(web.GlobalControllerRouter["openvpn-server-manage/controllers:APISignalController"],
			web.ControllerComments{
			Method: "Send",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

			web.GlobalControllerRouter["openvpn-server-manage/controllers:APISysloadController"] = append(web.GlobalControllerRouter["openvpn-server-manage/controllers:APISysloadController"],
			web.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

			web.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"] = append(web.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"],
			web.ControllerComments{
			Method: "Download",
			Router: `/certificates/:key`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

			web.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"] = append(web.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"],
			web.ControllerComments{
			Method: "Get",
			Router: `/certificates`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

			web.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"] = append(web.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"],
			web.ControllerComments{
			Method: "Post",
			Router: `/certificates`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}

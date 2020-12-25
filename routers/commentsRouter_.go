package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["joyconn-server-manage/controllers/base:APISessionController"] = append(beego.GlobalControllerRouter["joyconn-server-manage/controllers/base:APISessionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["joyconn-server-manage/controllers/base:APISessionController"] = append(beego.GlobalControllerRouter["joyconn-server-manage/controllers/base:APISessionController"],
        beego.ControllerComments{
            Method: "Kill",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["joyconn-server-manage/controllers/base:APISignalController"] = append(beego.GlobalControllerRouter["joyconn-server-manage/controllers/base:APISignalController"],
        beego.ControllerComments{
            Method: "Send",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["joyconn-server-manage/controllers/base:APISysloadController"] = append(beego.GlobalControllerRouter["joyconn-server-manage/controllers/base:APISysloadController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["joyconn-server-manage/controllers/vpn:CertificatesController"] = append(beego.GlobalControllerRouter["joyconn-server-manage/controllers/vpn:CertificatesController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/certificates",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["joyconn-server-manage/controllers/vpn:CertificatesController"] = append(beego.GlobalControllerRouter["joyconn-server-manage/controllers/vpn:CertificatesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/certificates",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["joyconn-server-manage/controllers/vpn:CertificatesController"] = append(beego.GlobalControllerRouter["joyconn-server-manage/controllers/vpn:CertificatesController"],
        beego.ControllerComments{
            Method: "Download",
            Router: "/certificates/:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

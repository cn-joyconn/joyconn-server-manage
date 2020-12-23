package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["openvpn-server-manage/controllers:APISessionController"] = append(beego.GlobalControllerRouter["openvpn-server-manage/controllers:APISessionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openvpn-server-manage/controllers:APISessionController"] = append(beego.GlobalControllerRouter["openvpn-server-manage/controllers:APISessionController"],
        beego.ControllerComments{
            Method: "Kill",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openvpn-server-manage/controllers:APISignalController"] = append(beego.GlobalControllerRouter["openvpn-server-manage/controllers:APISignalController"],
        beego.ControllerComments{
            Method: "Send",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openvpn-server-manage/controllers:APISysloadController"] = append(beego.GlobalControllerRouter["openvpn-server-manage/controllers:APISysloadController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/certificates",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/certificates",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["openvpn-server-manage/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Download",
            Router: "/certificates/:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

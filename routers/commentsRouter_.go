package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego_learning/controllers:CaptchaController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:CaptchaController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/get",
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_learning/controllers:CaptchaController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:CaptchaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/verify",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_learning/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserAdd",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_learning/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserDelete",
            Router: "/delete/:uid",
            AllowHTTPMethods: []string{"DELETE"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_learning/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserDelete2",
            Router: "/delete/batch",
            AllowHTTPMethods: []string{"DELETE"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_learning/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserList",
            Router: "/list",
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_learning/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserList2",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_learning/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_learning/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_learning/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_learning/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserUpdate",
            Router: "/update",
            AllowHTTPMethods: []string{"PUT"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

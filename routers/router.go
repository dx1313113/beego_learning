// Package routers @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"beego_learning/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 定义路由组 使用 Namespace func 来定义
	// 路由组的使用：
	// 在最外层 一般新建一个路由组 关键字：NewNamespace
	// 返回值用于调用 AddNamespace func 进行注册
	ns := beego.NewNamespace("/v1",
		// 在内部再次需要定义子路由组的时候。 可使用
		// NS 级别的Namespace 去定义 理论上Namespace是可以无限进行嵌套的
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/captcha",
			beego.NSInclude(
				&controllers.CaptchaController{},
			),
		),
	)
	//注册路由组
	beego.AddNamespace(ns)
}

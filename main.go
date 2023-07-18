package main

import (
	"beego_learning/global"
	"beego_learning/initialize"
	_ "beego_learning/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//初始化数据库
	global.BeeDb = initialize.OrmMysql()
	//初始化日志
	global.BeeLog = initialize.Log()
	//设置静态地址
	beego.SetStaticPath("/static", "static")

	beego.Run()
	//bee run -gendoc=true -downdoc=true
}

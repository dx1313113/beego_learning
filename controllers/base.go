package controllers

import (
	"beego_learning/models"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

type BaseController struct {
	web.Controller
	i18n.Locale
	user    models.User
	isLogin bool
}

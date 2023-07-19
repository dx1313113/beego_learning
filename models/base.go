package models

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"time"
)

type BaseController struct {
	web.Controller
	i18n.Locale
	user    User
	isLogin bool
}

type BaseModel struct {
	CreatedAt *time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt *time.Time `json:"updated_at" orm:"auto_now"`
	DeletedAt *time.Time `json:"deleted_at" `
}

type BaseDelete struct {
	IDS []int `json:"ids"`
}

type Response struct {
	Code int
	Msg  string
	Data any `json:"data" default:"nil"` //data结构体类型
	//Error error `json:"error" default:"nil"`
}

type Page struct {
	PageCurrent int
	PageSize    int
	TotalPage   int
	TotalCount  int64
	FirstPage   bool
	LastPage    bool
}

type Pagination struct {
	PageCurrent int
	PageSize    int
}

type PaginateResponse struct {
	Code int
	Msg  string
	Data any //data结构体类型
	Page Page
	//Error error `json:"error" default:"nil"`
}

func (c *BaseController) Respond(code int, message string, data ...interface{}) {
	c.Ctx.Output.SetStatus(code)
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	c.Data["json"] = struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}{
		Code:    code,
		Message: message,
		Data:    d,
	}
	c.ServeJSON()
}

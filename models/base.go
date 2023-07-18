package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
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

type PaginateResponse struct {
	Code int
	Msg  string
	Data any //data结构体类型
	Page Page
	//Error error `json:"error" default:"nil"`
}

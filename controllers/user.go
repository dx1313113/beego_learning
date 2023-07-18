package controllers

import (
	"beego_learning/models"
	"beego_learning/service"
	"encoding/json"
)

type UserControlller struct {
	BaseController
}

// 这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。
func (this *UserControlller) Prepare() {
	//this.user.Pwd = utils.ScryptPwd(this.user.Pwd)
	//this.StopRun()  //终止运行
}

// 这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
// 如果注册了就会通过interface来进行执行函数，性能上会提升很多
func (this *UserControlller) URLMapping() {
	this.Mapping("UserList", this.UserList)
	this.Mapping("UserAdd", this.UserAdd)
	this.Mapping("UserUpdate", this.UserUpdate)
	this.Mapping("UserDelete", this.UserDelete)
}

// UserList @Title UserList
// @Description Get all user list
// Param name query string true "查询条件"
// Param pageCurrent query int true "页码"
// Param pageSize query int true "页面大小"
// Param page body
// @router /list [Get]
func (this *UserControlller) UserList() {
	var req = make(map[string]any)
	req["name"] = this.GetString("name")
	req["pageCurrent"], _ = this.GetInt("pageCurrent", 1)
	req["pageSize"], _ = this.GetInt("pageSize", 10)
	if ret, err := service.UserList(req); err == nil {
		this.Data["json"] = models.Response{Code: 200, Msg: "查询成功", Data: ret}
	} else {
		this.Data["json"] = models.Response{Code: 500, Msg: "查询失败"}
	}
	this.ServeJSON()
}

// UserAdd @Title UserAdd
// @Description 新增事件
// @Param data body models.UserAdd true "添加用户"
// @Success 200 {int} models.User.Id
// @router /add [post]
func (this *UserControlller) UserAdd() {
	var user models.User
	var err error
	tt := this.Ctx.Input.RequestBody
	print(tt)
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		if ret, err := service.UserAdd(user); err == nil {
			this.Data["json"] = models.Response{Code: 200, Msg: "添加成功", Data: ret}
		} else {
			this.Data["json"] = models.Response{Code: 500, Msg: "添加失败"}
		}
	} else {
		this.Data["json"] = models.Response{Code: 500, Msg: "json格式解析失败"}
	}
	this.ServeJSON()
}

// UserUpdate @Title User_update
// @Description 更新事件
// @Param data body models.UserUpdate true "更新用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @router /update [PUT]
func (this *UserControlller) UserUpdate() {
	var user models.UserUpdate
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		if ret, err := service.UserUpdate(&user); err == nil {
			this.Data["json"] = models.Response{Code: 200, Msg: "修改成功", Data: ret}
		} else {
			this.Data["json"] = models.Response{Code: 500, Msg: "修改失败"}
		}
	}
	this.ServeJSON()
}

// UserDelete @Title User_delete
// @Description 批量删除事件
// @Param uid path string true "用户id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @router /delete/:uid [DELETE]
func (this *UserControlller) UserDelete() {
	var uid, _ = this.GetInt(":uid")
	if ret, err := service.UserDelete(uid); err == nil {
		this.Data["json"] = models.Response{Code: 200, Msg: "删除成功", Data: ret}
	} else {
		this.Data["json"] = models.Response{Code: 500, Msg: "删除失败"}
	}
	this.ServeJSON()
}

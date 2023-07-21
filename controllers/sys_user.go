package controllers

import (
	"beego_learning/middleware"
	"beego_learning/models"
	"beego_learning/service"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"net/http"
)

type UserController struct {
	models.BaseController
}

// 这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。
func (this *UserController) Prepare() {
	//this.user.Pwd = utils.ScryptPwd(this.user.Pwd)
	//this.StopRun()  //终止运行
}

// 这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
// 如果注册了就会通过interface来进行执行函数，性能上会提升很多
func (this *UserController) URLMapping() {
	this.Mapping("UserList", this.UserList)
	this.Mapping("UserList2", this.UserList2)
	this.Mapping("UserAdd", this.UserAdd)
	this.Mapping("UserUpdate", this.UserUpdate)
	this.Mapping("UserDelete", this.UserDelete)
	this.Mapping("UserDeleteBatch", this.UserDeleteBatch)
	this.Mapping("Login", this.Login)
	this.Mapping("Register", this.Register)
}

// UserList @Title UserList
// @Description 使用get获取用户信息
// Param name query string true "查询条件"
// Param pageCurrent query int true "页码"
// Param pageSize query int true "页面大小"
// @router /list [Get]
func (this *UserController) UserList() {
	var req = make(map[string]any)
	req["name"] = this.GetString("name")
	req["pageCurrent"], _ = this.GetInt("pageCurrent", 1)
	req["pageSize"], _ = this.GetInt("pageSize", 10)
	if ret, err := service.UserList(req); err != nil {
		this.Respond(http.StatusBadRequest, err.Error())
		return
	} else {
		this.Respond(http.StatusOK, "获取用户列表成功", ret)
	}
}

// UserList2 @Title UserList2
// @Description 使用post获取用户信息
// @Param data body models.UserQuery true "查询条件"
// @Success 200 {any} any
// @router /list [post]
func (this *UserController) UserList2() {
	var user models.UserQuery
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		if ret, err := service.UserList2(user); err != nil {
			this.Respond(http.StatusBadRequest, err.Error())
			return
		} else {
			this.Respond(http.StatusOK, "获取用户列表成功", ret)
		}
	}
}

// UserAdd @Title UserAdd
// @Description 新增事件
// @Param data body models.UserAdd true "添加用户"
// @Success 200 {int} models.User.Id
// @router /add [post]
func (this *UserController) UserAdd() {
	var user models.User
	var err error
	tt := this.Ctx.Input.RequestBody
	print(tt)
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		if ret, err := service.UserAdd(user); err != nil {
			this.Respond(http.StatusBadRequest, err.Error())
			return
		} else {
			this.Respond(http.StatusOK, "获取用户列表成功", ret)
		}
	} else {
		this.Respond(http.StatusBadRequest, "json格式解析失败")
	}
}

// UserUpdate @Title User_update
// @Description 更新用户
// @Param data body models.UserUpdate true "更新用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @router /update [PUT]
func (this *UserController) UserUpdate() {
	var user models.UserUpdate
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		if ret, err := service.UserUpdate(&user); err == nil {
			this.Data["json"] = models.Response{Code: http.StatusOK, Msg: "修改成功", Data: ret}
		} else {
			this.Data["json"] = models.Response{Code: http.StatusNotModified, Msg: "修改失败"}
		}
	}
	this.ServeJSON()
}

// UserDelete @Title User_delete
// @Description 删除用户
// @Param uid path string true "用户id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @router /delete/:uid [DELETE]
func (this *UserController) UserDelete() {
	var uid, _ = this.GetInt(":uid")
	if ret, err := service.UserDelete(uid); err == nil {
		this.Data["json"] = models.Response{Code: http.StatusOK, Msg: "删除成功", Data: ret}
	} else {
		this.Data["json"] = models.Response{Code: http.StatusNotModified, Msg: err.Error()}
	}
	this.ServeJSON()
}

// UserDeleteBatch @Title User_delete2
// @Description 批量删除用户
// @Param ids body models.BaseDelete true "用户ids"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @router /delete/batch [DELETE]
func (this *UserController) UserDeleteBatch() {
	var ids models.BaseDelete
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &ids); err == nil {
		if ret, err := service.UserDeleteBatch(ids); err == nil {
			this.Data["json"] = models.Response{Code: http.StatusOK, Msg: "删除成功", Data: ret}
		} else {
			this.Data["json"] = models.Response{Code: http.StatusNotModified, Msg: err.Error()}
		}
	} else {
		this.Data["json"] = models.Response{Code: http.StatusBadRequest, Msg: "json格式解析失败"}
	}
	this.ServeJSON()
}

// Login @Title Login
// @Description 用户登录验证
// @Param data body models.LoginRequest true "用户登录"
// @Success 200 {Respond} Respond
// @router /login
func (this *UserController) Login() {
	lr := new(models.LoginRequest)
	if store.Verify(lr.CaptchaId, lr.Captcha, true) {
		if err := this.unmarshalPayload(lr); err != nil {
			this.Respond(http.StatusBadRequest, err.Error())
			return
		}
		if lrs, statusCode, err := service.Login(lr); err != nil {
			this.Respond(statusCode, err.Error())
			return
		} else {
			next, err := middleware.TokenNext(lrs, 0)
			if err != nil {
				this.Respond(http.StatusBadRequest, err.Error())
				return
			}
			// 将token设置到Header
			this.Ctx.Output.Header("Authorization", next.Token)
			this.Respond(http.StatusOK, "登陆成功", next)

		}
	} else {
		this.Respond(400, "验证码错误")
	}
}

// Register
// @Title Register
// @Description 用户注册
// @Param data body models.RegisterRequest true "用户注册"
// @Success 200 {Respond} Respond
// @router /register [post]
func (this *UserController) Register() {
	cu := new(models.RegisterRequest)
	// 获取request body
	if err := this.unmarshalPayload(cu); err != nil {
		this.Respond(http.StatusBadRequest, err.Error())
	}
	createUser, statusCode, err := service.Register(cu)
	if err != nil {
		this.Respond(statusCode, err.Error())
		return
	}
	this.Respond(http.StatusOK, "", createUser)
}

// 解析请求，并将请求体存储到v中
// unmarshalPayload
// @Param	v	interface{}	true	"接收解析后的请求体的变量"
func (this *UserController) unmarshalPayload(v interface{}) error {
	// json 解析
	// Unmarshal(data []byte, v interface{})
	// 将json字符串解码到相应的数据结构
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if err != nil {
		logs.Error("RequestBody 解析失败！")
	}
	if err != nil {
		logs.Error("unmarshal payload of %s error: %s", this.Ctx.Request.URL.Path, err)
	}
	return nil
}

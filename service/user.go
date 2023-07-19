package service

import (
	"beego_learning/middleware"
	"beego_learning/models"
	"beego_learning/utils"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"net/http"
)

func UserAdd(req models.User) (id int64, err error) {
	req.Salt, err = utils.GenerateSalt()
	if err != nil {
		return id, errors.New("生成盐失败")
	}
	req.Pwd, err = utils.ScryptPwd(req.Pwd, req.Salt)
	if err != nil {
		return id, errors.New("加密失败")
	}
	o := orm.NewOrm()
	id, err = o.Insert(&req)
	if err == nil {
		return id, nil
	} else {
		return 0, errors.New("插入失败")
	}

}

func UserList2(req models.UserQuery) (any, error) {
	var users []*models.User
	var ret = make(map[string]any)
	var page models.Page
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user")
	if req.RealName != "" {
		qs = qs.Filter("real_name__icontains", req.RealName)
	}
	if req.UserName != "" {
		qs = qs.Filter("user_name__icontains", req.UserName)
	}
	total, _ := qs.Count()
	_, err := qs.OrderBy("-profile__id").Limit(req.Page.PageSize, req.Page.PageSize*(req.Page.PageCurrent-1)).All(&users)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	page = utils.PageUtil(total, req.Page.PageCurrent, req.Page.PageSize)
	ret["page"] = page
	ret["list"] = users
	return ret, nil
}

func UserList(req map[string]any) (any, error) {
	var users []*models.User
	var ret = make(map[string]any)
	var page models.Page
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user").Filter("name__icontains", req["name"]).OrderBy("-profile__id")
	total, _ := qs.Count()
	_, err := qs.Limit(req["pageSize"], req["pageSize"].(int)*(req["pageCurrent"].(int)-1)).All(&users)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	page = utils.PageUtil(total, req["pageCurrent"].(int), req["pageSize"].(int))
	ret["page"] = page
	ret["list"] = users
	return ret, nil
}

func UserUpdate(req *models.UserUpdate) (int64, error) {
	o := orm.NewOrm()
	var user = models.User{
		RealName: req.RealName,
	}
	user.Id = req.ID
	ret, err := o.Update(&user, "real_name")
	if err == nil {
		return ret, nil
	} else {
		return 0, err
	}
}

func UserDelete(req int) (any, error) {
	o := orm.NewOrm()
	n, err := o.Delete(&models.User{Id: req})
	if err == nil {
		return n, nil
	} else {
		return 0, errors.New("删除失败")
	}

}

func UserDelete2(req models.BaseDelete) (any, error) {
	o := orm.NewOrm()
	n, err := o.QueryTable("sys_user").Filter("id__in", req.IDS).Delete()
	if err == nil {
		return n, nil
	} else {
		return 0, errors.New("删除失败")
	}

}

func UserExist(req models.UserExist) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("sys_user").Filter("id", req.UserName).Filter("pwd", req.Pwd).Exist()
	return exist

}

// Login handles login request
func Login(lr *models.LoginRequest) (*models.LoginResponse, int, error) {
	//lr := new(models.LoginRequest)
	//获取用户名密码
	username := lr.UserName
	password := lr.Pwd
	// 验证用户名和密码是否为空
	if len(username) == 0 || len(password) == 0 {
		return nil,
			http.StatusBadRequest,
			errors.New("error: 用户名或密码为空")
	}
	o := orm.NewOrm()
	user := &models.User{UserName: username}
	err := o.Read(user, "id")
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("账号不存在")
	}
	hash, err := utils.ScryptPwd(password, user.Salt)
	if err != nil {
		return nil,
			http.StatusBadRequest, // 400
			err
	}
	if hash != user.Pwd {
		return nil,
			http.StatusBadRequest,
			errors.New("密码错误")
	}

	// 生成token
	tokenString, err := middleware.GenerateToken(user.Id, 0)
	if err != nil {
		return nil,
			http.StatusBadRequest,
			err
	}
	// 生成的token 返回给前端
	return &models.LoginResponse{
		UserName: user.UserName,
		Token:    tokenString,
	}, http.StatusOK, nil

}

// CreateUser creates a user
func Register(cr *models.RegisterRequest) (*models.RegisterResponse, int, error) {
	o := orm.NewOrm()
	// 检查用户名是否存在
	userNameCheck := models.User{UserName: cr.UserName}
	err := o.Read(&userNameCheck, "username")
	if err == nil {
		return nil, http.StatusBadRequest, errors.New("username has already existed")
	}
	// 生成 用户的加密的钥匙
	saltKey, err := utils.GenerateSalt()
	if err != nil {
		logs.Info(err.Error())
		return nil, http.StatusBadRequest, err
	}
	// 生成hash加密的密码
	hash, err := utils.ScryptPwd(cr.Pwd, saltKey)
	if err != nil {
		logs.Info(err.Error())
		return nil, http.StatusBadRequest, err
	}

	//创建用户
	user := models.User{
		UserName: cr.UserName,
		Pwd:      hash,
		Salt:     saltKey,
	}
	_, err = o.Insert(&user)
	if err != nil {
		logs.Info(err.Error())
		return nil, http.StatusBadRequest, err
	}

	return &models.RegisterResponse{
		Id:       user.Id,
		UserName: user.UserName,
	}, http.StatusOK, nil

}

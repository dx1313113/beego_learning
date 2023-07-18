package service

import (
	"beego_learning/models"
	"beego_learning/utils"
	"errors"
	"github.com/beego/beego/v2/client/orm"
)

func UserAdd(req models.User) (int64, error) {
	req.Pwd = utils.BcryptPwd(req.Pwd)
	o := orm.NewOrm()
	id, err := o.Insert(&req)
	if err == nil {
		return id, nil
	} else {
		return 0, errors.New("插入失败")
	}

}

func UserList(req map[string]any) (any, error) {
	var users []*models.User
	var ret = make(map[string]any)
	var page models.Page
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user").Filter("name__icontains", req["name"])
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
		Name: req.Name,
		Pwd:  utils.BcryptPwd(req.Pwd),
	}
	user.Id = req.ID
	ret, err := o.Update(&user, "Name", "Pwd")
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

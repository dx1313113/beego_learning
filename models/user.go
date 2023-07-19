package models

type User struct {
	BaseModel
	Id       int    `json:"id"`
	UserName string `json:"user_name" valid:"Required"`
	RealName string `json:"real_name"`
	Pwd      string `json:"-" valid:"Required"`
	Email    string `json:"email" valid:"Email;MaxSize(100)"`
	Mobile   string `json:"mobile" valid:"Mobile"`
	Salt     string `json:"-"`
}

// 自定义表名
// 自定义表名
func (u *User) TableName() string {
	return "sys_user"
}

type UserAdd struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"-" valid:"Required"`
	Email    string `json:"email" valid:"Email;MaxSize(100)"`
	Mobile   string `json:"mobile" valid:"Mobile"`
}

type UserUpdate struct {
	ID       int    `json:"id"`
	RealName string `json:"real_name"`
	Pwd      string `json:"-"`
	Email    string `json:"email" valid:"Email;MaxSize(100)"`
	Mobile   string `json:"mobile" valid:"Mobile"`
}

type UserQuery struct {
	UserName string     `json:"user_name"`
	RealName string     `json:"real_name"`
	Page     Pagination `json:"page"`
}

type UserExist struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type LoginRequest struct {
	UserName string `json:"user_name" valid:"required"`
	Pwd      string `json:"-" valid:"required"`
}

// LoginResponse 定义登录响应
type LoginResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

type RegisterRequest struct {
	UserName string `json:"user_name"  valid:"required"`
	Pwd      string `json:"-" valid:"required"`
	Salt     string `json:"-"`
}

// LoginResponse 定义登录响应
type RegisterResponse struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

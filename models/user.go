package models

type User struct {
	BaseModel
	Id   int    `json:"id"`
	Name string `json:"name" valid:"Required"`
	Pwd  string `json:"pwd" valid:"Required"`
}

// 自定义表名
func (u *User) TableName() string {
	return "sys_user"
}

type UserAdd struct {
	Name string `json:"name" valid:"Required"`
	Pwd  string `json:"pwd" valid:"Required"`
}

type UserUpdate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type UserQuery struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

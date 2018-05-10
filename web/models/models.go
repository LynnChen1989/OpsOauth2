package models

import (
	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	Id           int
	UserName     string //用户名
	UserPassword string //用户密码
}

type AppInfo struct {
	Id          int
	AppName     string                    //应用名称
	AppDesc     string                    //应用描述
	LogoPath    string                    //Logo路径
	CallbackUrl string                    //回调地址
	UserInfo    *UserInfo `orm:"rel(fk)"` //所属用户
}

func init() {
	orm.RegisterModel(new(UserInfo), new(AppInfo))
}

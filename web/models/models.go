package models

import "github.com/astaxie/beego/orm"

type UserInfo struct {
	Id           int
	UserName     string
	UserPassword string
}

type AppInfo struct {
	Id          int
	AppName     string
	AppDesc     string
	CallbackUrl string
	User        *UserInfo `orm:"rel(one)"`
}

func init() {
	orm.RegisterModel(new(UserInfo), new(AppInfo))
}

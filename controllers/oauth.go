package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"OpsOauth2/models"
)

type MainController struct {
	beego.Controller
}

func (mc *MainController) Get() {
	o := orm.NewOrm()
	appInfo:= new(models.AppInfo)
	qs := o.QueryTable(appInfo)
	var app []*models.AppInfo
	qs.All(&app)

	mc.Data["Apps"] = app
	mc.LayoutSections = make(map[string]string)
	mc.LayoutSections["Content"] = "application/applist.html"
	mc.Layout = "base.html"
	mc.TplName = "base.html"
}

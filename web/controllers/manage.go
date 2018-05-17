package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"web/models"
)

type ManageAppController struct {
	beego.Controller
}

func (ma *ManageAppController) Get() {

	o := orm.NewOrm()
	appInfo:= new(models.AppInfo)
	qs := o.QueryTable(appInfo)
	var app []*models.AppInfo
	qs.All(&app)

	ma.Data["Apps"] = app
	ma.LayoutSections = make(map[string]string)
	ma.LayoutSections["Content"] = "application/manage.html"
	ma.Layout = "base.html"
	ma.TplName = "base.html"
}

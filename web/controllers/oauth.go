package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"web/models"
	"log"
)

type MainController struct {
	beego.Controller
}

func (mc *MainController) Get() {
	o := orm.NewOrm()
	app := models.AppInfo{}
	err := o.Read(&app)

	mc.LayoutSections = make(map[string]string)
	mc.LayoutSections["Content"] = "application/applist.html"
	mc.Layout = "base.html"
	mc.TplName = "base.html"
}

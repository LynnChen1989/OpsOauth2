package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"OpsOauth2/models"
	"strconv"
)

type AddedAppController struct {
	beego.Controller
}

func (aac *AddedAppController) Get() {
	aac.LayoutSections = make(map[string]string)
	aac.LayoutSections["Content"] = "application/appadded.html"
	aac.Layout = "base.html"
	aac.TplName = "base.html"

	id := aac.Ctx.Input.Param(":id")
	Aid, _ := strconv.Atoi(id)
	app := models.AppInfo{Id: Aid}
	o := orm.NewOrm()
	o.Read(&app)
	aac.Data["app"] = app
}

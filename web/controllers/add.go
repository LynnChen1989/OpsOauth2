package controllers

import "github.com/astaxie/beego"

type AddAppController struct {
	beego.Controller
}

func (aac *AddAppController) Get() {
	aac.LayoutSections = make(map[string]string)
	aac.LayoutSections["Content"] = "application/appadd.html"
	aac.Layout = "base.html"
	aac.TplName = "base.html"
}

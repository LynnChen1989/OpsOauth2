package controllers

import (
	"github.com/astaxie/beego"
)

type AddedAppController struct {
	beego.Controller
}

func (aac *AddedAppController) Get() {
	aac.LayoutSections = make(map[string]string)
	aac.LayoutSections["Content"] = "application/appadded.html"
	aac.Layout = "base.html"
	aac.TplName = "base.html"
}

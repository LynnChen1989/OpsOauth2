package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type AddAppController struct {
	beego.Controller
}

func (aac *AddAppController) Get() {
	aac.LayoutSections = make(map[string]string)
	aac.LayoutSections["Content"] = "application/appadd.html"
	aac.Layout = "base.html"
	aac.TplName = "base.html"
}

func (aac *AddAppController) Post() {
	appName := aac.Input().Get("app-name")
	appDesc := aac.Input().Get("app-description")
	appLogo := aac.Input().Get("app-logo")
	appCallBack := aac.Input().Get("app-callback")
	fmt.Println(appName, appDesc, appLogo, appCallBack)
	aac.Redirect("/app-added", 302)
}

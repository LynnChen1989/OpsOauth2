package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Content"] = "applist.html"
	//c.LayoutSections["AppTest"] = "apptest.html"
	//c.Layout = "applist.html"
	c.TplName = "base.html"
}

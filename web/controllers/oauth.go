package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"web/models"
	"fmt"
	"log"
)

type MainController struct {
	beego.Controller
}

func (mc *MainController) Get() {
	//mc.Data["Website"] = "beego.me"
	//mc.Data["Email"] = "astaxie@gmail.com"
	o := orm.NewOrm()
	user := models.UserInfo{Id: 1}
	err := o.Read(&user)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user.UserName)
	mc.LayoutSections = make(map[string]string)
	mc.LayoutSections["Content"] = "application/applist.html"
	mc.Layout = "base.html"
	mc.TplName = "base.html"
}

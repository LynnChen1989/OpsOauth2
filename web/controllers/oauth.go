package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"web/models"
	"fmt"
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
	if err == orm.ErrNoRows {
		fmt.Println("Can not find any data.")
	} else if err == orm.ErrMissPK {
		fmt.Println("Can not find pk.")
	} else {
		fmt.Println(user.UserName)
	}
	mc.LayoutSections = make(map[string]string)
	mc.LayoutSections["Content"] = "application/applist.html"
	mc.Layout = "base.html"
	mc.TplName = "base.html"
}

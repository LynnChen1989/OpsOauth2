package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"log"
	"strings"
	"github.com/astaxie/beego/orm"
	"web/models"
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
	appCallBack := aac.Input().Get("app-callback")
	saveStatus := aac.SavePic(appName, "app-logo")

	if saveStatus == true {
		logoPath := fmt.Sprintf("static/img/upload/%s.png", appName)
		id, _ := aac.SaveData(appName, appDesc, logoPath, appCallBack)
		aac.Redirect(fmt.Sprintf("/app-added/%d", id), 302)
	} else {
		aac.Redirect("/app-add", 302)
	}

}

// 保存图片
func (aac *AddAppController) SavePic(appName, domLogo string) (status bool) {
	// appName：应用名称；
	// domLogo：logo html 的input标签name属性；
	file, header, err := aac.GetFile(domLogo)
	if err != nil {
		log.Println("Handle logo error:", err)
	}
	initFileName := header.Filename
	suffix := strings.Split(initFileName, ".")[1]
	if suffix != "png" {
		log.Println("[E] Just support png picture.")
		status = false
		return
	}
	logoPath := fmt.Sprintf("static/img/upload/%s.png", appName)
	if file != nil {
		err := aac.SaveToFile("app-logo", logoPath)
		if err != nil {
			log.Println("Save logo picture error:", err)
			status = false
			return
		} else {
			log.Println("Save logo picture success.")
			status = true
			return
		}
	}
	return
}

// 保存数据库记录
func (aac *AddAppController) SaveData(appName, appDesc, logoPath, callbackUrl string) (id int64, err error) {
	o := orm.NewOrm()
	var app models.AppInfo
	user := models.UserInfo{UserName: "chenlin002"}
	errRead := o.Read(&user, "UserName")
	if errRead == orm.ErrNoRows {
		log.Println("[E] can not find user.")
	} else if errRead == orm.ErrMissPK {
		log.Println("[E] can not find pk")
	} else {
		log.Println("[I] relation user is:", user.UserName)
	}

	app.AppName = appName
	app.AppDesc = appDesc
	app.LogoPath = logoPath
	app.CallbackUrl = callbackUrl
	app.UserInfo = &user

	id, err = o.Insert(&app)
	if err != nil {
		log.Println("[E] Insert data error:", err)
	} else {
		log.Println("[I] Insert data success, id is:", id)
	}
	return
}

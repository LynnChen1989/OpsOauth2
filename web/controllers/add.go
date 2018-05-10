package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"log"
	"strings"
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

	aac.SavePic(appName, "app-logo")
	fmt.Println(appName, appDesc, appCallBack)
	aac.Redirect("/app-added", 302)
}

// 保存图片
func (aac *AddAppController) SavePic(appName string, domLogo string) {
	// appName：应用名称；
	// domLogo：logo html 的input标签name属性；
	file, header, err := aac.GetFile(domLogo)
	if err != nil {
		log.Println("Handle logo error:", err)
	}
	initFileName := header.Filename
	picSuffix := strings.Split(initFileName, ".")[1]
	newFileName := fmt.Sprintf("%s.%s", appName, picSuffix)
	logoPath := fmt.Sprintf("static/img/upload/%s", newFileName)
	if file != nil {
		err := aac.SaveToFile("app-logo", logoPath)
		if err != nil {
			log.Println("Save logo picture error:", err)
		} else {
			log.Println("Save logo picture success.")
		}
	}
}

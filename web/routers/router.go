package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/app-add", &controllers.AddAppController{})
	beego.Router("/app-added", &controllers.AddedAppController{})
}

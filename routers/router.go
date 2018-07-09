package routers

import (
	"OpsOauth2/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/app-add", &controllers.AddAppController{})
	beego.Router("/app-added/?:id([0-9]+)", &controllers.AddedAppController{})
	beego.Router("/app-manage", &controllers.ManageAppController{})
}

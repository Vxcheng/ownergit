package routers

import (
	"github.com/astaxie/beego"
	"ownergit/beego_new/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

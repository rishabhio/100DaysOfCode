package routers

import (
	"selflove/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/love", &controllers.LoveController{})
}

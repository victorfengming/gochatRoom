package routers

import (
	"github.com/astaxie/beego"
	"github.com/victorfengming/gochatroom/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/chatRoom", &controllers.ServerController{})
	beego.Router("/chatRoom/WS", &controllers.ServerController{}, "get:WS")
}

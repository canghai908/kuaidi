package routers

import (
	"github.com/astaxie/beego"
	"github.com/canghai908/kuaidi/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/ucenter/", &controllers.UcenterController{})
	beego.Router("/api/", &controllers.UserController{}, "get:Info")
	beego.Router("/user/signin/", &controllers.UserController{}, "get:Register")
}

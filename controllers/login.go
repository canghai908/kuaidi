package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/canghai908/kuaidi/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// 判断是否为退出操作
	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}

	this.TplNames = "login.html"
}
func (this *LoginController) Post() {
	// 获取表单信息

	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	// 验证用户名及密码
	uinfo, _ := models.GetUser(uname)

	if uname == uinfo.Username &&
		pwd == uinfo.Password {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}

	this.Redirect("/ucenter/", 302)
	return

}
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value
	uinfo, _ := models.GetUser(uname)
	return uname == uinfo.Username &&
		pwd == uinfo.Password

}

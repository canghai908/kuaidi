package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	"github.com/canghai908/kuaidi/models"
)

type UcenterController struct {
	beego.Controller
}

func (this *UcenterController) Get() {
	// 判断是否为退出操作
	this.Data["IsHome"] = true
	this.TplNames = "ucenter.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	models.Init()

	if checkAccount(this.Ctx) {
		p := this.Ctx.GetCookie("uname")
		t, _ := models.GetUser(p)
		this.Data["Username"] = t.Username
		this.Data["Key"] = t.Key
		this.Data["Count"] = t.Count
	} else {
		this.Data["Err"] = "你还没有登录！"
	}
}

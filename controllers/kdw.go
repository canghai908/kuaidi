package controllers

import (
	"github.com/astaxie/beego"
	"github.com/canghai908/kuaidi/models"
)

type UserController struct {
	beego.Controller
}

type Message struct {
	Message interface{}
	RetCode int64
}

func (this *UserController) Register() {
	models.Init()
	message := Message{}
	uname := this.GetString("uname")
	pwd := this.GetString("pwd")
	if uname == "" || pwd == "" {
		message.Message = "用户名和密码为必填项"
		message.RetCode = 1
	} else {
		err := models.AddUser(uname, pwd)
		if err != nil {
			message.Message = "成功"
			message.RetCode = 0
		} else {

			message.Message = "服务器内部处理错误"
			message.RetCode = 999
		}
	}

	this.Data["json"] = message
	this.ServeJson()
}

func (this *UserController) Info() {
	models.Init()
	message := Message{}
	uname := this.GetString("unm")
	key := this.GetString("key")
	cemskin := this.GetString("cemskind")
	cno := this.GetString("cno")
	if key == "" || cemskin == "" || cno == "" {
		message.Message = "授权key，快递类型和快递号为必填项！"
		message.RetCode = 1
	} else {
		p, err := models.GetUser(uname)
		if err != nil {
			message.Message = "用户未注册！"
			message.RetCode = 1
		} else {
			kid := p.Key
			if key != kid {
				message.Message = "用户名与key不匹配！"
				message.RetCode = 1
			}
		}

	}

	this.Data["json"] = message
	this.ServeJson()

}

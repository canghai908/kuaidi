package main

import (
	"github.com/astaxie/beego"
	"github.com/canghai908/kuaidi/models"
	_ "github.com/canghai908/kuaidi/routers"
)

func main() {
	models.Init()
	beego.Run()
}

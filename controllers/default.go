package controllers

import (
	//"github.com/astaxie/beego"
	//beego "github.com/astaxie/beego/context"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//c.Ctx.WriteString("helloa")

}

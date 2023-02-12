package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Hello() {
	c.Data["Title"] = "Hello Beego!"
	c.Data["SubTitle"] = "First Beego Route!"
	c.TplName = "hello.tpl"
}

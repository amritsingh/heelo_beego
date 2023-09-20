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
	// Get the query parameters from the URL
	name := c.GetString("name", "")
	age, _ := c.GetInt("age", 0)
	c.Data["Name"] = name
	c.Data["Age"] = age
	c.TplName = "hello.tpl"
}

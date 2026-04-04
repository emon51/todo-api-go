package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RootController struct {
	beego.Controller
}

func (c *RootController) Get() {
	c.Data["json"] = map[string]string{
		"message": "Hello, World",
	}

	c.ServeJSON()
}
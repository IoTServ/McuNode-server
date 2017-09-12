package controllers

import (
	"github.com/astaxie/beego"
	_"strings"
)

type TermController struct {
	beego.Controller
}

func (c *TermController) Get() {
	id := c.Ctx.Input.Param(":id")

	c.Data["id"] = id
	c.TplName = "terminal.html"
}

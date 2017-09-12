package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"

	//data, err := Asset("views/index.html")
	//if err != nil {
	//	// Asset was not found.
	//}
	//c.Ctx.WriteString(string(data))
}

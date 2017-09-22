package main

import (
	_ "mcunode/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	beego.BConfig.AppName = "McuNode"
	beego.BConfig.Listen.HTTPPort=80
	beego.BConfig.RunMode="prod"
	beego.BConfig.ServerName = "McuNode Server"
	beego.BConfig.EnableErrorsShow = false
	beego.BConfig.EnableErrorsRender=false
	beego.BConfig.WebConfig.FlashName = "MCUNODE_FLASH"
	beego.BConfig.WebConfig.Session.SessionName = "MCUNODEsessionID"
	beego.Run()
	fmt.Printf("Please open this web server port : 80")
}

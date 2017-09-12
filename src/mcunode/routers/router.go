package routers

import (
	"mcunode/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/term/:id:string", &controllers.TermController{})
	beego.Router("/ws/:id:string", &controllers.WebSocketController{}, "get:Join")
	beego.Router("/proxy/:id:string/*", &controllers.ProxyController{})
	//beego.Router("/*", &controllers.StaticController{})
}

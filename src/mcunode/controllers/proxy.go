package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"strings"
)

type ProxyController struct {
	beego.Controller
}

func (c *ProxyController) Get() {
	id := c.Ctx.Input.Param(":id")
	//file := c.Ctx.Input.Param(":splat")
	//c.Ctx.Input.Params()
	URI:=c.Ctx.Request.RequestURI
	//fmt.Printf(URI)
	url:=strings.Replace(URI, "/proxy/"+id+"/", "", -1)
	//fmt.Printf(url)
	//c.Data["id"] = id
	//c.Data["file"] = file
	delete(httpmsg,id)
	msg:=""
	wspublish <- wsEvent(id, []byte(url))
	//c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	if strings.Contains(url, ".css") {
		c.Ctx.Output.Header("Content-Type", "text/css; charset=utf-8")
	}else if strings.Contains(url, ".js") {
		c.Ctx.Output.Header("Content-Type", "application/x-javascript; charset=utf-8")
	}else {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}
	time.Sleep(time.Millisecond*250)
	if v, ok := httpmsg[id]; ok {
		msg=v
		delete(httpmsg,id)
		//fmt.Printf("start")
		c.Ctx.WriteString(msg)
		return
	}else {
		time.Sleep(time.Millisecond*100)
		if v, ok := httpmsg[id]; ok {
			msg = v
			delete(httpmsg, id)
			//fmt.Printf("start")
			c.Ctx.WriteString(msg)
			return
		}else{
			time.Sleep(time.Millisecond*300)
			if v, ok := httpmsg[id]; ok {
				msg = v
				delete(httpmsg, id)
				//fmt.Printf("start")
				c.Ctx.WriteString(msg)
				return
			}else{
				c.Ctx.WriteString("TimeOut")
				return
			}
		}
	}

}


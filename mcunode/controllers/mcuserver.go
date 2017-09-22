package controllers

import (
	"github.com/gorilla/websocket"

	"mcunode/models"
	"net"
	"runtime"
	"fmt"
	"strings"
	"time"
)

func wsEvent(id string, msg []byte) models.Event {
	return models.Event{id, msg}
}

func mcuEvent(id string, msg []byte) models.Event {
	return models.Event{id, msg}
}

func Join(user string, ws *websocket.Conn) {
	fmt.Printf("Join")
	subscribe <- Subscriber{Id: user, Conn: ws}
}
type Subscriber struct {
	Id string
	Conn *websocket.Conn // Only for WebSocket users; otherwise nil.
}

var (
	mcu map[string]net.Conn
	ws map[string]*websocket.Conn
	httpmsg map[string]string
	// Send events here to publish them.
	wspublish = make(chan models.Event, 100)
	mcupublish = make(chan models.Event, 100)
	subscribe = make(chan Subscriber, 100)
)

// This function handles all incoming chan messages.
func mcuhandler(id string,conn net.Conn) {
	fmt.Printf("mcu-id:"+id)
	//conn.SetReadDeadline(time.Now().Add(time.Second*6))
	//conn.SetWriteDeadline(time.Now().Add(time.Second*3))
	if v, ok := mcu[id]; ok {
		v.Close()
		delete(mcu,id)
	}
	mcu[id]=conn
	for {
		var recv []byte = make([]byte, 10240)
		if conn.SetReadDeadline(time.Now().Add(time.Minute*60))!=nil{
			conn.Close()
			fmt.Printf("设置超时退出！")
			runtime.Goexit()
		}
		n,e:=conn.Read(recv)
		if e!=nil{
			conn.Close()
			//	delete(mcu,id)
			fmt.Printf("读取错误退出！"+e.Error())
			runtime.Goexit()
		}
		if conn.SetReadDeadline(time.Time{}) !=nil{
			conn.Close()
			fmt.Printf("取消超时退出！")
			runtime.Goexit()
		}
		msg:=recv[:n]
		if string(msg)!="" && string(msg)!="<h1></h1>" {
			str := strings.Replace(string(msg), "<h1></h1>", "", -1)
			mcupublish <- mcuEvent(id, []byte(str))
		}
	}
}

func wshandler(newws Subscriber) {
	// Message receive loop.
	id :=newws.Id
	wsconn :=newws.Conn
	fmt.Printf("ws-id:"+id)
	if v, ok := ws[id]; ok {
		v.Close()
		delete(ws,id)
	}
	ws[id]=wsconn
	for {
		if wsconn.SetReadDeadline(time.Now().Add(time.Minute*60))!=nil{
			wsconn.Close()
			fmt.Printf("设置超时退出！")
			runtime.Goexit()
		}
		_, p, err := wsconn.ReadMessage()
		if err != nil {
			wsconn.Close()
			fmt.Printf("WS读取超时退出！！")
			runtime.Goexit()
		}
		if wsconn.SetReadDeadline(time.Time{}) !=nil{
			wsconn.Close()
			fmt.Printf("WS取消超时退出！")
			runtime.Goexit()
		}
		wspublish <- wsEvent(id, p)
	}
}

func runmcuhandler(l net.Listener)  {
	for {	clientconnn, _ := l.Accept()
		var recv []byte = make([]byte, 10240)
		n, _ := clientconnn.Read(recv)
		id := string(recv[:n])
		go mcuhandler(id, clientconnn)
	}
}

func runwshandler()  {
	for {
		select {
		case newws := <-subscribe:
			go wshandler(newws)
		}
	}
}

func bridge()  {
	for {
		select {
			//case <-time.After(time.Second * 2): <-mcupublish
			case mcumsg := <-mcupublish:
				id:=mcumsg.User
				msg:=mcumsg.Content
				fmt.Printf("ws开始处理")
				if v, ok := ws[id]; ok {
					fmt.Println(id+"ws存在,并发送到WS")
					error:=v.WriteMessage(1,msg)
					if error !=nil {
						v.Close()
						//delete(ws,id)
						//mcupublish<-mcumsg
						httpmsg[id] = httpmsg[id]+string(msg)
					}

				} else {
					//mcupublish<-mcumsg
					httpmsg[id] = httpmsg[id]+string(msg)
					fmt.Println("WS Key Not Found")
				}

			case wsmsg := <-wspublish:
				fmt.Printf("wsmsg收到消息")
				id:=wsmsg.User
				msg:=wsmsg.Content
				if v, ok := mcu[id]; ok {
					if string(msg)!="" {
						fmt.Println(id + "mcu存在"+string(msg))
						_,error:=v.Write(msg)
						if error!=nil{
							v.Close()
							fmt.Printf(id+"设备离线")
						}
					}
				}else {fmt.Println("MCU Key Not Found")}
			}

	}
}

func init() {
	ws = make(map[string]*websocket.Conn)
	mcu = make(map[string]net.Conn)
	httpmsg = make(map[string]string)
	fmt.Printf("MCUNODE (McuNode.com) server started!\n")
	runtime.GOMAXPROCS(runtime.NumCPU())
	l, _ := net.Listen("tcp", ":8001")
	go runmcuhandler(l)
	go runwshandler()
	go bridge()
}

package bcsocket

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/ValdekGreen/bcdns/bcfile"
	"net/http"
	"strings"
)

var names *bcfile.Parser

//This function only giveaways a information of Name object that stated in the URL
func StrServer(ws *websocket.Conn) {
	fmt.Println(ws.Request().URL.String())
	answ, err := names.Names[strings.Split(ws.Request().URL.String(), "?")[1]].ReadBytesArmored()
	if err != nil {
		panic(err)
	}
	websocket.Message.Send(ws, answ)
}

//This function handls full upd messages communication
func UpdServer(ws *websocket.Conn) {
	var msg string
	websocket.Message.Receive(ws, &msg)
	fmt.Println(msg)
}

func FileServerHandler(p *bcfile.Parser) {
	names = p
	http.Handle("/bcget", websocket.Handler(StrServer))
	http.Handle("/bcupd", websocket.Handler(UpdServer))
	err := http.ListenAndServe(":8054", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

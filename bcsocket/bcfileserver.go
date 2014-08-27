package bcsocket

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/ValdekGreen/bcdns/bcfile"
	"net/http"
	"strings"
)

var names *bcfile.Parser

func StrServer(ws *websocket.Conn) {
	fmt.Println(ws.Request().URL.String())
	answ, err := names.Names[strings.Split(ws.Request().URL.String(), "?")[1]].ReadBytesArmored()
	if err != nil {
		panic(err)
	}
	websocket.Message.Send(ws, answ)
}

func FileServerHandler(p *bcfile.Parser) {
	names = p
	http.Handle("/bcget", websocket.Handler(StrServer))
	err := http.ListenAndServe(":8054", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

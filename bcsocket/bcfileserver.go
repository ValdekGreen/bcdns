package bcsocket

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/ValdekGreen/bcdns/bcfile"
	"io"
	"net/http"
	"os"
	"strings"
)

func Server(ws *websocket.Conn) {
	fmt.Println(ws.Request().URL.String())
	io.Copy(ws, ws)
}

func NewFilesRequestHandler(z *bcfile.Zone) func(ws *websocket.Conn) {
	str, err := z.ReadStringArmored()
	if err != nil {
		panic(err)
	}
	strr := strings.NewReader(str)
	return func(ws *websocket.Conn) {
		io.Copy(ws, strr)
	}
}

func generate_handlers(root string) {
	zone, err := os.Open(root)
	defer zone.Close()
	if err != nil {
		panic(err)
	}
	http.Handle("bc/"+root, NewFilesRequestHandler(new(Zone)))
}

func FileServerHandler() {
	generate_handlers("testicullo/")
	http.Handle("/bc", websocket.Handler(Server))
	err := http.ListenAndServe(":8054", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

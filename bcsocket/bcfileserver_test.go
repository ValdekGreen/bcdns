package bcsocket

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"testing"
)

func TestStrServer(t *testing.T) {
	origin := "ws://localhost:8054/bcget?valdek.tzone"
	url := "ws://localhost:8054/bcget?admin.valdek.tzone"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg string
	fmt.Println(websocket.Message.Receive(ws, &msg))
	fmt.Printf("Received: %s.\n", msg)
}

// func TestUpdServer(t *testing.T) {
// 	origin := "ws://localhost:8054/bcupd?valdek.tzone"
// 	url := "ws://localhost:8054/bcupd?valdek.tzone"
// 	ws, err := websocket.Dial(url, "", origin)
// 	websocket.Message.Send(ws, "")
// }

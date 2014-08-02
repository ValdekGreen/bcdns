package bcsocket

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"testing"
)

func TestEcho(t *testing.T) {
	origin := "http://localhost/"
	url := "ws://localhost:8054/bc"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}

package bcsocket

import (
	"code.google.com/p/go.net/websocket"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

func TestHashing(t *testing.T) {
	origin := "ws://localhost:8054/bcget?home.valdek.tzone"
	url := "ws://localhost:8054/bcget?admin.valdek.tzone"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
	}
	//Here we casting the package who requests a admin.valdek.tzone updates comparing with hash and sends it to home.valdek.tzone
	p := CastPackage(ws, "home.valdek.tzone", TypeUpdReq)
	fmt.Println("The Hashing package: ", p.String())
	h := md5.New()
	endp_contents, err := ioutil.ReadFile("testicullo/tzone/valdek/admin.endp")
	if err != nil {
		panic(err)
	}
	io.WriteString(h, string(endp_contents))
	resulting_hash := h.Sum(nil)
	if UnmarshalPackage(p.String()).body != "?"+string(resulting_hash) {
		fmt.Println("ERROR in TestHashing: hashes are not respective")
		t.Fail()
	}
	p.conn.Write(p.Bytes())
}

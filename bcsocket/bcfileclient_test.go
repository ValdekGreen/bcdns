package bcsocket

import (
	"code.google.com/p/go.net/websocket"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

func TestHashingEndpoint(t *testing.T) {
	origin := "ws://localhost:8054/bcget?home.valdek.tzone"
	url := "ws://localhost:8054/bcget?admin.valdek.tzone"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
	}
	//Here we casting the package who requests a admin.valdek.tzone updates comparing with hash and sends it to home.valdek.tzone
	p := CastPackage(ws, "home.valdek.tzone", TypeUpdReq)
	h := md5.New()
	endp_contents, err := ioutil.ReadFile("testicullo/tzone/valdek/admin.endp")
	if err != nil {
		panic(err)
	}
	io.WriteString(h, string(endp_contents))
	resulting_hash := h.Sum(nil)
	if string([]byte(p.String())[38:]) != string(resulting_hash) {
		fmt.Println("ERROR in TestHashing: hashes are not respective")
		fmt.Println(string(resulting_hash) + " vs " + string([]byte(p.String())[38:])) //a length of package header + ? byte
		t.Fail()
	}
	p.conn.Write(p.Bytes())
}

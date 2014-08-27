package bcfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var root_own = new(owner)
var vald_own = new(owner)

var tzone = new(Zone)         //.tzone
var xxxzx = new(Zone)         //.xxxzx.tzone
var mytunelspace = new(Zone)  //.valdek.tzone~valdek.xxxzx.tzone
var adminsite = new(endpoint) //admin.valdek.tzone
var homepc = new(endpoint)
var awesomesite = new(endpoint) //boobies.tzone

func TestInit(t *testing.T) {
	root_own.ReadKeyRing("root", []byte("passroot"))
	vald_own.ReadKeyRing("vald", []byte("passvald"))
	tzone.New(nil, "tzone", root_own)
	xxxzx.New(tzone, "xxxzx", root_own)
	mytunelspace.New(tzone, "valdek", root_own) //will be delegated
	adminsite.New(mytunelspace, "admin", root_own)
	homepc.New(mytunelspace, "home", root_own)
	awesomesite.New(tzone, "boobies", root_own)
}

func TestName2Path(t *testing.T) {
	za := new(Zone)
	zb := new(Zone)
	za.New(nil, "hype", root_own)
	zb.New(za, "test", root_own)
	if zb.Path() != root+"hype/test/" {
		t.Fail()
	}
	os.RemoveAll(zb.Path())
}

func TestWrite(t *testing.T) {
	//r_tzone := []string{"tzone	3600	IN	AAAA	::1", "tzone	3600	IN	A	127.0.0.1"}
	r_tzone := []string{"AAAA	::1", "A	127.0.0.1"}
	r_mts_admin := []string{"AAAA 	fc0e:69cb:d79c:97f9:0e1b:1656:2b3f:2829", "A 	192.168.0.1"}
	r_mts_homepc := []string{"AAAA 	fc1e:69cb:d59c:97f9:0e1b:1256:2b3f:2211", "A 	192.168.0.2"}
	root_own.records = make(map[Name][]string)
	root_own.records[tzone] = r_tzone
	root_own.Write(tzone)
	root_own.records[adminsite] = r_mts_admin
	root_own.Write(adminsite)
	root_own.records[homepc] = r_mts_homepc
	root_own.Write(homepc)
}

func TestZoneReader(t *testing.T) {
	be, err := mytunelspace.ReadString()
	if err != nil {
		panic("Read zone " + mytunelspace.FullName() + " unsuccesfull" + be)
	}
}

func TestDelegate(t *testing.T) {
	mytunelspace.delegate(vald_own)
	fcheck, err := os.Open(mytunelspace.Path() + "../;DGATE")
	defer fcheck.Close()
	if err != nil {
		defer t.Fail()
		panic(err)
	}
	b := make([]byte, 1000)
	fcheck.Read(b)
	signm := append([]byte("valdek:"), []byte(vald_own.label)...)
	if !strings.Contains(string(b), string(signm)) {
		t.Fail()
	}
}

func TestSigner(t *testing.T) {
	sigm, _ := mytunelspace.ReadBytesArmored()
	fmt.Print("Signing now: " + string(sigm))
	vald_own.Sign(mytunelspace)
	defer func(z *Zone) {
		if r := recover(); r != nil {
			fmt.Println("Recovered::TestSigner", r)
			sigf, _ := os.Open(mytunelspace.Path() + ";SIG")
			defer sigf.Close()
			signature, _ := ioutil.ReadAll(sigf)
			fmt.Println(signature)
			t.Fail()
		}
	}(mytunelspace)
	vald_own.Check(mytunelspace)
}

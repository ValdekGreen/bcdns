package bcfile

import (
	"os"
	"strings"
	"testing"
)

var root_own = new(owner)
var vald_own = new(owner)

var tzone = new(zone)         //.tzone
var xxxzx = new(zone)         //.xxxzx.tzone
var mytunelspace = new(zone)  //.valdek.tzone~valdek.xxxzx.tzone
var adminsite = new(endpoint) //admin.valdek.tzone
var homepc = new(endpoint)
var awesomesite = new(endpoint) //boobies.tzone

func TestInit(t *testing.T) {
	root_own.ReadKeyRing("root")
	vald_own.ReadKeyRing("vald")
	tzone.new(nil, "tzone", root_own)
	xxxzx.new(tzone, "xxxzx", root_own)
	mytunelspace.new(tzone, "valdek", root_own) //will be delegated
	adminsite.new(mytunelspace, "admin", root_own)
	homepc.new(mytunelspace, "home", root_own)
	awesomesite.new(tzone, "boobies", root_own)
}

func TestName2Path(t *testing.T) {
	za := new(zone)
	zb := new(zone)
	za.new(nil, "hype", root_own)
	zb.new(za, "test", root_own)
	if zb.Path() != root+"hype/test/" {
		t.Fail()
	}
	os.RemoveAll(zb.Path())
}

func TestWrite(t *testing.T) {
	r_tzone := []string{"AAAA ::1", "A 127.0.0.1"}
	r_mts_admin := []string{"AAAA fc0e:69cb:d79c:97f9:0e1b:1656:2b3f:2829", "A 192.168.0.1"}
	r_mts_homepc := []string{"AAAA fc1e:69cb:d59c:97f9:0e1b:1256:2b3f:2211", "A 192.168.0.2"}
	root_own.records = make(map[name][]string)
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

func TestSigner(t *testing.T) {
	root_own.Sign(mytunelspace)
}

func TestDelegate(t *testing.T) {
	mytunelspace.delegate(vald_own)
	fcheck, err := os.Open(mytunelspace.Path() + "../DGATE")
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

func TestSignature(t *testing.T) {

}

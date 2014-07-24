package bcfile

import (
	pgp "code.google.com/p/go.crypto/openpgp"
	"os"
	"testing"
)

var root_own = new(owner)
var rootp, _ = os.Open("keys/root.priv")

var vald_own = new(owner)

var tzone = new(zone)         //.tzone
var xxxzx = new(zone)         //.xxxzx.tzone
var adminsite = new(endpoint) //admin.valdek.tzone
var homepc = new(endpoint)
var awesomesite = new(endpoint) //boobies.tzone
var mytunelspace = new(zone)    //.valdek.tzone~valdek.xxxzx.tzone

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

func TestInit(t *testing.T) {
	owns, _ := pgp.ReadKeyRing(rootp)
	root_own.own = *owns[0]
	tzone.new(nil, "tzone", root_own)
	xxxzx.new(tzone, "xxxzx", root_own)
	awesomesite.new(tzone, "boobies", root_own)
	adminsite.new(mytunelspace, "admin", root_own) //will be delegated
	homepc.new(mytunelspace, "home", root_own)
	mytunelspace.new(tzone, "valdek", root_own)
}

func TestWrite(t *testing.T) {
	r_tzone := []string{"AAAA ::1", "A 127.0.0.1"}
	r_mts_admin := []string{"AAAA fc0e:69cb:d79c:97f9:0e1b:1656:2b3f:2869", "A 192.168.0.1"}
	root_own.records = make(map[name][]string)
	root_own.records[tzone] = r_tzone
	root_own.Write(tzone)
	root_own.records[adminsite] = r_mts_admin
	root_own.Write(adminsite)
}

func TestZoneReader(t *testing) {
	zr := ZoneReader{mytunelspace}
}

package bcutil

import (
	//"os"
	"testing"
)

var tzone = new(zone)           //.tzone
var xxxzx = new(zone)           //.xxxzx.tzone
var adminsite = new(endpoint)   //valdek.xxxzx.tzone
var awesomesite = new(endpoint) //boobies.tzone
var mytunelspace = new(zone)    //.valdek.tzone~valdek.xxxzx.tzone

func TestInit(t *testing.T) {
	root = "testicullo/"
	tzone.new(nil, "tzone")
	xxxzx.new(tzone, "xxxzx")
	awesomesite.new(tzone, "boobies")
	adminsite.new(xxxzx, "valdek")
	mytunelspace.new(tzone, "valdek")
}

func TestTree(t *testing.T) {
	root = "~/go/tests"
	mytunelspace.delegate(adminsite.in.origin)
}

package bcutil

import (
	"net"
)

var root string = "" //Root path, need to set externally

//A directory type
type zone struct {
	name   string
	in     *zone //<nil> means \\
	origin *endpoint
}

//A file type -- the concrette pc
type endpoint struct {
	addr net.IP //address
	name string //awesomepc.xxx.zzz.hype| where awesomepc is name
	in   *zone  //zone, can be <nil> for root domains
}

//interface to implement zone and endpoint types
type name interface {
	new(in *zone) name     //new file or folder | if in==nil -> file/folder will be created in contect root dir
	move(to *zone)         //local
	move_admin(to *zone)   //move remote
	delegate(to *endpoint) //bring admin rights
	ZonesNames() string    //get everything exept name
	FullName() string      //get everything
	Path() string          //get path relative to name root
}

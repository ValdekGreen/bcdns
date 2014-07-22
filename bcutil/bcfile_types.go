package bcutil

import (
	"github.com/miekg/dns"
	"net"
)

var root string = "" //Root path, need to set externally

//A directory type
type zone struct {
	name   string
	in     *zone //<nil> means \\
	origin *owner
}

type owner struct { //the fields recorded to a file
	pub_key string //it will be another type someday
	records *[]dns.RR
}

//A file type -- the concrette pc
type endpoint struct {
	addr net.IP //address
	name string //awesomepc.xxx.zzz.hype| where awesomepc is name
	in   *zone  //zone, can be <nil> for root domains
}

//interface to implement zone and endpoint types
type name interface {
	//blank --becouse there can be only addr field in it
	new(in *zone, nm string) //new file or folder | if in==nil -> file/folder will be created in contect root dir
	move(to *zone)           //local
	move_admin(to *zone)     //move remote
	delegate(to *owner)      //change the data
	ZonesNames() string      //get everything exept name
	FullName() string        //get everything
	Path() string            //get path relative to name root
}

func _inline_ohno(err error) {
	if err != nil {
		panic(err)
	}
}

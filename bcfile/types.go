package bcfile

import (
	pgp "code.google.com/p/go.crypto/openpgp"
)

var root string = "testicullo/" //Root path, need to set externally

//A directory type
type zone struct {
	name   string
	in     *zone  //<nil> means \\
	origin *owner //it recorded to endpoint in same thone with same name
}

type owner struct { //the fields recorded to a file
	own     pgp.Entity
	records map[name][]string //maps names to records
}

//A file type -- the concrette pc
type endpoint struct {
	origin *owner //address
	name   string //awesomepc.xxx.zzz.hype| where awesomepc is name
	in     *zone  //zone, can be <nil> for root domains
}

//interface to implement zone and endpoint types
type name interface {
	//new method must be called after memory
	new(in *zone, nm string, own *owner) //new file or folder | if in==nil -> file/folder will be created in contect root dir
	move(to *zone)                       //local
	move_admin(to *zone)                 //move remote
	delegate(to *owner)                  //change the data
	ZonesNames() string                  //get everything exept name
	Name() string                        //get name only
	FullName() string                    //get everything
	Path() string                        //get path relative to name root
	Owner() *owner                       //get owner
}

func _inline_ohno(err error) {
	if err != nil {
		panic(err)
	}
}

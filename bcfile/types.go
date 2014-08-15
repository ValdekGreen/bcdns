package bcfile

import (
	pgp "code.google.com/p/go.crypto/openpgp"
)

var root string = "testicullo/" //Root path, need to set externally

//A directory type
type Zone struct {
	name   string
	in     *Zone  //<nil> means \\
	origin *owner //it recorded to endpoint in same thone with same Name
	endp   map[string]*endpoint
}

type owner struct { //the fields recorded to a file
	label   string //the label that is used in a keys/
	own     pgp.Entity
	records map[Name][]string //maps Names to records
}

//A file type -- the concrette pc
type endpoint struct {
	origin *owner //address
	name   string //awesomepc.xxx.zzz.hype| where awesomepc is Name
	in     *Zone  //zone, can be <nil> for root domains
}

//interface to implement zone and endpoint types
type Name interface {
	//new method must be called after memory
	New(in *Zone, nm string, own *owner) //new file or folder | if in==nil -> file/folder will be created in contect root dir
	move(to *Zone)                       //local
	move_admin(to *Zone)                 //move remote
	delegate(to *owner)                  //change the data
	ZonesNames() string                  //get everything exept Name
	Name() string                        //get Name only
	FullName() string                    //get everything
	Path() string                        //get path relative to Name root
	Owner() *owner                       //get owner
	ReadStringArmored() (str string, err error)
}

func _inline_ohno(err error) {
	if err != nil {
		panic(err)
	}
}

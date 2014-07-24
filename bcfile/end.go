package bcfile

import (
	"os"
)

func (e *endpoint) new(in *zone, nm string, own *owner) {
	e.in = in
	e.name = nm
	var err error = nil
	path := ""
	if in == nil {
		path = root
	} else {
		path = e.Path()
	}
	f, err := os.Create(path + e.name)
	defer f.Close()
	f.Chmod(0700)
	/*
		!TODO: Infile authority and metadata writing
	*/
	if err != nil {
		panic(err)
	}
}

func (e *endpoint) delegate(to *owner) {
	if e.in != nil {
		e.in.origin = to
	} /*else {
	!TODO: Infile authority and metadata
	} */
}

func (e *endpoint) move_admin(to *zone) {
	os.Rename(e.Path()+e.name, to.Path()+e.name)
}

func (e *endpoint) move(to *zone) {
	/*
		!TODO: Infile authority and metadata
	*/
}

func (e *endpoint) Name() string {
	return e.name
}

func (e *endpoint) ZonesNames() string {
	return e.in.FullName() //.name.yyy.zzz.hype -- example
}

func (e *endpoint) FullName() string {
	return e.name + e.ZonesNames() //end.name.yyy.zzz.hype -- example
}

func (e *endpoint) Path() string {
	return e.in.Path()
}

func (e *endpoint) Owner() *owner {
	return e.origin
}

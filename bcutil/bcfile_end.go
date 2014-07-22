package bcutil

import (
	"os"
)

func (e *endpoint) new(in *zone) name {
	e.in = in
	if in != nil {
		err := os.MkdirAll(root+e.Path()+"~", os.ModeDir)
		if err != nil {
			panic(err)
		}
		return e
	}
	err := os.MkdirAll(root+e.name, os.ModeDir)
	if err != nil {
		panic(err)
	}
	return e
}

func (e *endpoint) delegate(to *endpoint) {
	if e.in != nil {
		e.in.origin = to
	} /*else {
	!TODO: Infile authority and metadata
	} */
}

func (e *endpoint) move_admin(to *zone) {
	os.Rename(root+e.Path(), to.Path()+e.name)
}

func (e *endpoint) move(to *zone) {
	/*
		!TODO: Infile authority and metadata
	*/
}

func (e *endpoint) ZonesNames() string {
	return e.in.FullName() //.name.yyy.zzz.hype -- example
}

func (e *endpoint) FullName() string {
	return e.name + e.ZonesNames() //end.name.yyy.zzz.hype -- example
}

func (e *endpoint) Path() string {
	return e.in.Path() + "/" + e.name
}

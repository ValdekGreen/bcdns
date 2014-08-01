package bcfile

import (
	"os"
)

func (e *endpoint) new(in *zone, nm string, own *owner) {
	e.in = in
	if e.in == nil {
		panic("The endpoint" + nm + "isn't properly initializated: zone is ab")
	}
	e.name = nm
	in.endp[nm] = e
	/*
		!TODO: Infile authority and metadata writing
	*/
}

func (e *endpoint) delegate(to *owner) {
	if e.in != nil {
		e.in.origin = to
	}
	e.origin.AddName(e)
	e.in.delegate(to) //delegating the endpoint means delegating a full zone *.endpoint.xxx...
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
	if e.in == nil {
		panic("The endpoint" + e.Name() + "isn't properly initializated")
	}
	return e.in.Path()
}

func (e *endpoint) Owner() *owner {
	return e.origin
}

package bcutil

import (
	"os"
	"strings"
)

func (z *zone) new(in *zone, nm string) {
	z.in = in
	z.name = nm
	var err error = nil
	if in != nil {
		err = os.MkdirAll(root+z.Path(), 0700)
	} else {
		err = os.MkdirAll(root+z.name, 0700)
	}
	if err != nil {
		panic(err)
	}
}

func (z *zone) delegate(to *owner) {
	z.origin = to
}

func (z *zone) move_admin(to *zone) {
	os.Rename(root+z.Path(), to.Path()+z.name) //probably not works! test and implement a bicycle
}

func (z *zone) move(to *zone) {
	/*
		!TODO: Infile authority and metadata
	*/
}

func (z *zone) ZonesNames() string {
	s := ""
	for {
		if z.in != nil {
			s = s + "." + z.in.name
			z = z.in
		} else {
			return s //.yyy.zzz.hype -- example
		}
	}
}

func (z *zone) FullName() string {
	return "." + z.name + z.ZonesNames() //.name.yyy.zzz.hype -- example
}

func (z *zone) Path() string {
	inverse := func(strarr []string) []string {
		for i, j := 0, len(strarr)-1; i < j; i, j = i+1, j-1 {
			strarr[i], strarr[j] = strarr[j], strarr[i]
		}
		return strarr //["hype", "zzz", "yyy", "name"]
	}
	strarr := inverse(strings.Split(z.FullName(), ".")) //on inp: ".name.yyy.zzz.hype" -> ["name", "yyy", "zzz", "hype"]
	return strings.Join(strarr, "/")                    //"hype/zzz/yyy/name"
}

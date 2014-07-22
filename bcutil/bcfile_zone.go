package bcutil

import (
	"os"
	"strings"
)

func (z *zone) new(in *zone) name {
	z.in = in
	if in != nil {
		err := os.MkdirAll(root+z.Path()+"~"+z.origin.FullName(), os.ModeDir)
		if err != nil {
			panic(err)
		}
		return z
	}
	err := os.MkdirAll(root+z.name+"~"+z.origin.FullName(), os.ModeDir)
	if err != nil {
		panic(err)
	}
	return z
}

func (z *zone) delegate(to *endpoint) {
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
		strarrb := strarr
		k := 0
		for i := len(strarr); i > 0; i-- {
			strarrb[k] = strarr[i]
			k++
		}
		return strarrb //["hype", "zzz", "yyy", "name"]
	}
	strarr := inverse(strings.Split(z.FullName(), ".")) //on inp: ".name.yyy.zzz.hype" -> ["name", "yyy", "zzz", "hype"]
	return strings.Join(strarr, "/")                    //"hype/zzz/yyy/name"
}

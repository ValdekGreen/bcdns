package bcfile

import (
	"os"
	"strings"
)

func (z *Zone) New(in *Zone, nm string, own *owner) {
	z.in = in
	z.name = nm
	var err error = nil
	if in != nil {
		err = os.MkdirAll(z.Path(), 0700)
	} else {
		err = os.MkdirAll(root+z.name, 0700)
	}
	endf, _ := os.Open(z.Path())
	defer endf.Close()
	z.endp = make(map[string]*endpoint)
	if err != nil {
		panic(err)
	}
}

func (z *Zone) delegate(to *owner) {
	z.origin = to
	dgate, err := os.OpenFile(z.Path()+"../;DGATE", os.O_APPEND, 0666)
	if os.IsNotExist(err) {
		dgate, err = os.Create(z.Path() + "../;DGATE")
	}
	defer dgate.Close()
	if err != nil {
		panic(err.Error())
	}
	dgate.Write([]byte(z.Name() + ":" + to.label + "\n"))
	to.Sign(z)
}

func (z *Zone) move_admin(to *Zone) {
	os.Rename(root+z.Path(), to.Path()+z.name) //probably not works! test and implement a bicycle
}

func (z *Zone) move(to *Zone) {
	/*
		!TODO: Infile authority and metadata
	*/
}

func (z *Zone) Name() string {
	return z.name
}

func (z *Zone) ZonesNames() string {
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

func (z *Zone) FullName() string {
	return "." + z.name + z.ZonesNames() //.name.yyy.zzz.hype -- example
}

func (z *Zone) Path() string {
	inverse := func(strarr []string) []string {
		for i, j := 0, len(strarr)-1; i < j; i, j = i+1, j-1 {
			strarr[i], strarr[j] = strarr[j], strarr[i]
		}
		return strarr //["hype", "zzz", "yyy", "name"]
	}
	strarr := inverse(strings.Split(z.FullName(), ".")) //on inp: ".name.yyy.zzz.hype" -> ["name", "yyy", "zzz", "hype"]
	return root + strings.Join(strarr, "/")             //"hype/zzz/yyy/name"
}

func (z *Zone) Owner() *owner {
	return z.origin
}

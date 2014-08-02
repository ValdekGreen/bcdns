package bcfile

import (
	"fmt"
	"os"
	"strings"
)

type Parser struct {
	init_path string
	names     map[string]Name //all names found nil to init
}

//Type Zone is implements it too with [Path() string] method from Name interface and parse in this file
type NameParser interface {
	Path() string
	Parse(origin *Parser, from NameParser)
}

func correct_name(name string) bool {
	return name != "SIG" || name != "DGATE"
}

func FilterNames(names []string) []string {
	j := 0
	var filtered []string = make([]string, len(names))
	for i := 0; i < len(names); i++ {
		if names[i] != "SIG" || names[i] != "DGATE" {
			filtered[j] = names[i]
			j++
		}
	}
	return filtered[:j]
}

func parse_loop(origin *Parser, from NameParser) {
	var croot *os.File
	var err error
	if from != nil {
		croot, err = os.Open(from.Path())
	} else {
		croot, err = os.Open(origin.Path())
	}
	defer croot.Close()
	if err != nil {
		panic(err)
	}
	names, _ := croot.Readdirnames(0)
	for _, el := range names {
		// nmf, err := os.OpenFile(from.Path()+el, os.O_RDONLY, 0666)
		// defer nmf.Close()
		// if err != nil {
		// 	panic(err)
		// }
		if el == "DGATE" || el == "SIG" {
			continue
		}
		fmt.Println(names)
		fmt.Println(el)
		if strings.Contains(el, ".") { //then it is endpoint
			end := new(endpoint)
			end.New(from.(*Zone), strings.Split(el, ".")[0], nil) //own -- tbd
			origin.names[end.FullName()] = end
		} else {
			zn := new(Zone)
			if from != nil {
				zn.New(from.(*Zone), el, nil)
			} else {
				zn.New(nil, el, nil)
			}
			origin.names[zn.FullName()] = zn
			go zn.Parse(origin, zn)
		}
	}
}

//Finds all names in the path and records it to origin.endp;
//Can be called with any origin and from parameters
func (p *Parser) Parse(origin *Parser, from NameParser) {
	p.names = make(map[string]Name)
	parse_loop(p, nil)
}

func (z *Zone) Parse(origin *Parser, from NameParser) {
	parse_loop(origin, z)
}

func (p *Parser) Path() string {
	return p.init_path
}

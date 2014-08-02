package bcfile

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type Parser struct {
	sync.RWMutex
	init_path string
	names     map[string]Name //all names found nil to init
}

//Type Zone is implements it too with [Path() string] method from Name interface and Parse() in this file
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
		if !strings.Contains(el, ";") { //Filtering the service files
			if strings.Contains(el, ".") { //then it is endpoint
				end := new(endpoint)
				end.New(from.(*Zone), strings.Split(el, ".")[0], nil) //own -- tbd
				go func() {                                           //async getting an owner within a closure
					fps, err := os.OpenFile(end.Path()+end.Name()+".endp", os.O_RDONLY, 0666)
					defer fps.Close()
					if err != nil {
						panic(err)
					}
					read := bufio.NewReader(fps)
					//record, errr := read.ReadString('\n')
					own := new(owner)
					own.records = make(map[Name][]string)
					var str string = ""
					strs := []string{}
					for {

						str, err = read.ReadString('\n')
						if err != nil {
							if err == io.EOF {
								break
							}
							panic(err.Error() + " in " + end.Path() + end.Name() + ".endp")
						}
						strs = append(strs, str)
					}
					own.records[end] = strs
				}()
				log.Println("Found new endpoint " + end.Path() + end.Name() + ".endp")
				origin.Lock()
				origin.names[end.FullName()] = end
				origin.Unlock()
			} else {
				zn := new(Zone)
				if from != nil {
					zn.New(from.(*Zone), el, nil)
				} else {
					zn.New(nil, el, nil)
				}
				log.Println("Found new zone " + zn.Path())
				origin.Lock()
				origin.names[zn.FullName()] = zn
				origin.Unlock()
				zn.Parse(origin, zn)
			}
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
	parse_loop(origin, from)
}

func (p *Parser) Path() string {
	return p.init_path
}

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
	Init_path string
	Names     map[string]Name //all Names found nil to init
}

//Type Zone is implements it too with [Path() string] method from Name interface and Parse() in this file
type NameParser interface {
	Path() string
	Parse(origin *Parser, from NameParser)
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
	Names, _ := croot.Readdirnames(0)
	for _, el := range Names {
		own := new(owner)
		if !strings.Contains(el, ";") { //Filtering the service files
			if strings.Contains(el, ".") { //then it is endpoint
				end := new(endpoint)
				end.New(from.(*Zone), strings.Split(el, ".")[0], own) //own -- infunc eq
				go func() {                                           //async getting an owner within a closure
					fps, err := os.OpenFile(end.Path()+end.Name()+".endp", os.O_RDONLY, 0666)
					defer fps.Close()
					if err != nil {
						panic(err)
					}
					read := bufio.NewReader(fps)
					go func() {
						dgate, err := os.Open(end.Path() + "../;DGATE")
						defer dgate.Close()
						if err != nil {
							if os.IsNotExist(err) {
								log.Println(end.Path() + "../;DGATE is not exists")
								return
							} else {
								panic(err.Error() + " in reading dgate in " + end.Path() + "../;DGATE")
							}
						}
						bdgate := bufio.NewReader(dgate)
						labelstr, err := bdgate.ReadString(':')
						if err != nil {
							panic(err.Error() + " in reading dgate zonename of " + end.Name())
						}
						if labelstr == end.in.Name() {
							labelstr, err = bdgate.ReadString('\n')
							if err != nil {
								panic(err.Error() + " in reading dgate label of " + end.Name())
							}
						}
						//own.ReadKeyRing(labelstr)
					}()
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
				origin.Names[end.FullName()] = end
				origin.Unlock()
				end.origin = own
			} else {
				zn := new(Zone)
				if from != nil {
					zn.New(from.(*Zone), el, own)
				} else {
					zn.New(nil, el, own)
				}
				log.Println("Found new zone " + zn.Path())
				origin.Lock()
				origin.Names[zn.FullName()] = zn
				origin.Unlock()
				zn.Parse(origin, zn)
			}
		}
	}
}

//Finds all Names in the path and records it to origin.endp;
//Can be called with any origin and from parameters
func (p *Parser) Parse(origin *Parser, from NameParser) {
	p.Names = make(map[string]Name)
	parse_loop(p, nil)
}

func (z *Zone) Parse(origin *Parser, from NameParser) {
	parse_loop(origin, from)
}

func (p *Parser) Path() string {
	return p.Init_path
}

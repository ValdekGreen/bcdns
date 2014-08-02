package bcfile

import (
	"os"
	"strings"
)

//Just appends all zone endpoints contents into single string
func (z *Zone) ReadString() (str string, err error) {
	dir, eopenzone := os.Open(z.Path())
	if eopenzone != nil {
		panic("Error: " + z.FullName() + " can't read")
	}
	names, _ := dir.Readdirnames(0)
	var p []byte = make([]byte, len(names)*1024) //1 KB per file for records
	var b []byte = make([]byte, 1024)
	for i := 0; i < len(names); i++ {
		if !strings.Contains(names[i], ".endp") {
			continue
		}
		z.endp[strings.Split(names[i], ".")[0]].Read(b)
		p = append(p, b...)
	}
	return string(p), err
}

//Appends all zone endpoints contents into single string with armor:
//"endpointname:<contents>\n;"
func (z *Zone) ReadStringArmored() (str string, err error) {
	dir, eopenzone := os.Open(z.Path())
	if eopenzone != nil {
		panic("Error: " + z.FullName() + " can't read")
	}
	names, _ := dir.Readdirnames(0)
	var p []byte = make([]byte, len(names)*1024) //1 KB per file for records
	var b []byte = make([]byte, 1024)
	for i := 0; i < len(names); i++ {
		if !strings.Contains(names[i], ".endp") {
			continue
		}
		z.endp[strings.Split(names[i], ".")[0]].Read(b)
		p = append(p, []byte(z.endp[strings.Split(names[i], ".")[0]].Name()+":")...)
		p = append(p, b...)
		p = append(p, ';')
	}
	return string(p), err
}

func (e *endpoint) Read(p []byte) (n int, err error) {
	var efile *os.File
	efile, err = os.Open(e.Path() + e.name + ".endp")
	defer efile.Close()
	efile.Read(p)
	if err != nil {
		panic("Error reading:" + e.FullName() + " " + err.Error())
	}
	return
}

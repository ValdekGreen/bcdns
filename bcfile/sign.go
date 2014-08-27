package bcfile

import (
	"io/ioutil"
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
func (z *Zone) ReadBytesArmored() (byt []byte, err error) {
	dir, eopenzone := os.Open(z.Path())
	if eopenzone != nil {
		panic("Error: " + z.FullName() + " can't read")
	}
	names, _ := dir.Readdirnames(0)
	// var p []byte = make([]byte, len(names)*1024) //1 KB per file for records
	// var b []byte = make([]byte, 1024)
	var p []byte
	var b []byte
	for i := 0; i < len(names); i++ {
		if !strings.Contains(names[i], ".endp") {
			continue
		}
		b, err = z.endp[strings.Split(names[i], ".")[0]].ReadBytes()
		p = append(p, []byte(z.endp[strings.Split(names[i], ".")[0]].Name()+":")...)
		p = append(p, b...)
		p = append(p, ';')
	}
	return p, err
}

//2bImplemented
func (e *endpoint) ReadBytesArmored() (byt []byte, err error) {
	return e.ReadBytes()
}

func (e *endpoint) ReadBytes() (b []byte, err error) {
	efile, _ := os.Open(e.Path() + e.name + ".endp")
	return ioutil.ReadAll(efile)
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

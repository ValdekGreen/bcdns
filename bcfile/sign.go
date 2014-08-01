package bcfile

import (
	"os"
	"strings"
)

func (z *zone) ReadString() (str string, err error) {
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
		p = append(b, p...)
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

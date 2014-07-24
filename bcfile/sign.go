package bcfile

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type ZoneReader struct { //Reads all zones endpoints contents in one stream by sort.Strings()
	z zone
}

func (zr *ZoneReader) Read(p []byte) (n int, err error) {
	f, _ := os.Open(zr.z.Path())
	fmt.Println(f.Readdirnames(0))
	return
}

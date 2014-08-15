package bcfile

import (
	"fmt"
	"testing"
)

func TestParserTesticullo(t *testing.T) {
	p := new(Parser)
	p.Init_path = "testicullo/"
	p.Parse(nil, nil)
	for k, v := range p.Names {
		fmt.Println(k, v)
		fmt.Println(v.Owner())
	}
}

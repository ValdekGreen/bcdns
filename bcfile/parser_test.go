package bcfile

import (
	"fmt"
	"testing"
)

func TestParserTesticullo(t *testing.T) {
	p := new(Parser)
	p.init_path = "testicullo/"
	p.Parse(nil, nil)
	for k, v := range p.names {
		fmt.Println(k, v)
	}
}

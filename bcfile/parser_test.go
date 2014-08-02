package bcfile

import (
	"testing"
)

func TestParserTesticullo(t *testing.T) {
	p := Parser{"testicullo/", nil}
	p.Parse(nil, nil)
}

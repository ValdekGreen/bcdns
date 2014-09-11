package bcsocket

import (
	"code.google.com/p/go.net/websocket"
	"crypto/md5"
	"fmt"
	"github.com/ValdekGreen/bcdns/bcfile"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

//This const includes all type bytes
const (
	TypeUpdReq = 'U'
	TypeUpdAns = 'u'
	TypeStrSig = 'S'
	TypeStrZon = 's'
)

type PackageHeader struct {
	name    string
	emitter string
	time    []byte
	t       byte
}

type Package struct {
	header PackageHeader
	conn   *websocket.Conn
	body   string
}

func (p *Package) AddToBody(b []byte) {
	p.body = p.body + string(b)
}

//Create a new package with current time
func CastPackage(ws *websocket.Conn, em string, ty byte) *Package {
	p := new(Package)
	p.header.emitter = em
	p.conn = ws
	p.header.t = ty
	timencode, err := time.Now().UTC().MarshalBinary()
	if err != nil {
		panic(err)
	}
	p.header.time = timencode
	return p
}

func (p *Package) Bytes() []byte {
	return []byte(p.String())
}

//Just giveaway all contents with header
func (p *Package) String() string {
	var symbols string
	h := md5.New()
	switch p.header.t {
	case TypeUpdReq:
		symbols = "?"
		Name_name := strings.Split(p.conn.RemoteAddr().String(), "?")[1]
		if []byte(Name_name)[0] == '.' { //if first symbol of
			fmt.Println(Name_name + "It is zone")
			z := new(bcfile.Zone)
			z.New(nil, strings.Split(Name_name, ".")[1], nil)
			str, err := z.ReadString()
			if err != nil {
				panic(err)
			}
			fmt.Println("zone_content: " + str)
			io.WriteString(h, str)
		} else {
			endpoint_content, err := ioutil.ReadFile(bcfile.NameOfNameToPath(Name_name) + ".endp")
			if err != nil {
				panic(err)
			}
			io.WriteString(h, string(endpoint_content))
		}
		var err error = nil
		if err != nil {
			panic("*Package.String()::ReadStringArmored in Name " + err.Error())
		}
	default:
		symbols = ""
	}
	symbols = symbols + string(h.Sum(nil))
	p.body = symbols
	return strings.Join([]string{p.header.name,
		p.header.emitter,
		string(p.header.time),
		string(p.header.t)}, ":") + ":" + symbols //second ':' is for diff header from body
}

//Creates new package from string
func UnmarshalPackage(packoded string) *Package {
	p := new(Package)
	parsed := strings.Split(packoded, ":")
	p.header.emitter = parsed[0]
	p.header.time = []byte(parsed[1])
	p.header.t = []byte(parsed[2])[0]
	p.body = parsed[3]
	return p
}

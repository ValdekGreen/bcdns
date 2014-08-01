package bcfile

import (
	//"bufio"
	pgp "code.google.com/p/go.crypto/openpgp"
	//"fmt"
	"os"
	"strings"
)

func OwnerFromDGATE(z *zone) *owner {
	dgate, err := os.OpenFile(z.Path()+"../DGATE", os.O_RDONLY, 0666)
	defer dgate.Close()
	if err != nil {
		panic(err)
	}
	//bdgate := bufio.NewReader(dgate)
	//str, err := bdgate.ReadString('\n')
	if err != nil {
		panic(err)
	}
	own := new(owner)
	return own //!NOT WORKS
}

func (o *owner) ReadKeyRing(l string) {
	if l != "" {
		o.label = l
	}
	fpriv, _ := os.Open("keys/" + l + ".priv") //error not matters here cause in most cases it would not exist
	defer fpriv.Close()
	var el pgp.EntityList = *new(pgp.EntityList)
	el, _ = pgp.ReadKeyRing(fpriv)
	if len(el) > 0 {
		o.own.PrivateKey = el[0].PrivateKey
	}
}

func (o *owner) Sign(z *zone) {
	sigf, err := os.Create(z.Path() + "SIG")
	if err != nil {
		panic("SIG file from " + z.FullName() + " open failed " + err.Error())
	}
	defer sigf.Close()
	msg, errmsg := z.ReadString()
	if errmsg != nil {
		panic("Reading zone " + z.FullName() + " failed")
	}
	pgp.DetachSign(sigf, &o.own, strings.NewReader(msg), nil)
}

func (o *owner) Write(to name) {
	f, err := os.Create(to.Path() + to.Name() + ".endp")
	defer f.Close()
	for i := 0; i < len(o.records)-1; i++ {
		f.WriteString(o.records[to][i] + "\n")
	}
	if err != nil {
		panic(err)
	}
}

func (o *owner) AddName(what name) {
	o.records[what] = []string{}
}

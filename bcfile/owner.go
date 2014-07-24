package bcfile

import (
	pgp "code.google.com/p/go.crypto/openpgp"
	"fmt"
	//packet "code.google.com/p/go.crypto/openpgp/packet"
	"os"
)

func (o *owner) Sign(namefile string, sigfile string) {
	f, _ := os.Open(namefile)
	fsig, _ := os.Create(sigfile)
	defer f.Close()
	defer fsig.Close()
	pgp.DetachSign(fsig, &o.own, f, nil)
}

func (o *owner) Write(to name) {
	f, err := os.Create(to.Path() + to.Name() + ".endp")
	fmt.Println(to.Path() + to.Name() + ".endp" + "->" + to.Path() + to.Name() + ".endp.sig")
	defer f.Close()
	for i := 0; i < len(o.records); i++ {
		f.WriteString(o.records[to][i] + "\n")
	}
	o.Sign(to.Path()+to.Name()+".endp", to.Path()+to.Name()+".endp.sig")
	if err != nil {
		panic(err)
	}
}

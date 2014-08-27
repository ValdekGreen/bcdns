package bcfile

import (
	"errors"
	"github.com/kisom/cryptutils/common/public"
	"io/ioutil"
	"os"
)

//Pass can be nil.

func (o *owner) ReadKeyRing(l string, pass []byte) {
	if l != "" {
		o.label = l
	}
	o.own = new(Keypair)
	fpub, err := os.Open("keys/" + l + ".pub")
	defer fpub.Close()
	pub, err := ioutil.ReadAll(fpub)
	if err != nil {
		panic(err.Error() + ":Reading of public file")
	}
	pubk, err := public.UnmarshalPublic(pub)
	if err != nil {
		panic(err.Error() + ":Unmarshaling")
	}
	o.own.pub = pubk
	fpriv, err := os.Open("keys/" + l + ".priv")
	defer fpriv.Close()
	if os.IsNotExist(err) {
		return
	}
	p, err := ioutil.ReadAll(fpriv) //ignoring
	if err != nil {
		panic(err)
	}
	o.own.priv, err = public.ImportPrivate(p, pass)
	if err != nil {
		panic(err)
	}
}

func (o *owner) Sign(z *Zone) {
	// sigf, err := os.Open(z.Path() + ";SIG")
	// if err != nil {
	// 	panic(";SIG file from " + z.FullName() + " open failed " + err.Error())
	// }
	// defer sigf.Close()
	msg, errmsg := z.ReadBytesArmored()
	if errmsg != nil {
		panic("Reading zone " + z.FullName() + " failed")
	}
	if o.own.priv != nil {
		signature, ok := public.Sign(o.own.priv, msg)
		if !ok {
			panic("Something not ok in signing by " + o.label + " of " + z.FullName())
		}
		ioutil.WriteFile(z.Path()+";SIG", signature, 770)
		//sigf.Write(signature)
	} else {
		return
	}
}

func (o *owner) Write(to Name) {
	f, err := os.Create(to.Path() + to.Name() + ".endp")
	defer f.Close()
	for i := 0; i < len(o.records)-1; i++ {
		f.WriteString(o.records[to][i] + "\n")
	}
	if err != nil {
		panic(err)
	}
}

func (o *owner) AddName(what Name) {
	o.records[what] = []string{}
}

func (o *owner) Check(z *Zone) bool {
	sigf, err := os.Create(z.Path() + ";SIG")
	if err != nil {
		panic(";SIG file from " + z.FullName() + " open failed " + err.Error())
	}
	defer sigf.Close()
	signed, err := z.ReadBytesArmored()
	if err != nil {
		panic("Reading of " + z.FullName() + " failed " + err.Error())
	}
	sig, err := ioutil.ReadAll(sigf)
	ok := public.Verify(o.own.pub, []byte(signed), sig)
	if !ok {
		panic(errors.New(o.label + " invalid signature of " + z.FullName()))
	}
	return ok
}

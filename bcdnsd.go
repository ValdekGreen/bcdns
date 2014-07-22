package main

import (
	"flag"
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"
	"syscall"
)

var (
	printf   *bool
	compress *bool
	pool     *bool
	tsig     *string
)

func handleQ(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeAAAA:
		{
			rr := new(dns.AAAA)
			rr.AAAA = net.ParseIP("::1")
			m.Answer = []dns.RR{rr}
		}
	case dns.TypeA:
		{
			rr := new(dns.A)
			rr.A = net.ParseIP("127.0.0.1")
			m.Answer = []dns.RR{rr}
		}
	case dns.TypeTXT:
		{
			rr := new(dns.TXT)
			rr.Txt = []string{"OK"}
			m.Answer = []dns.RR{rr}
		}
	}
	fmt.Print(m.String())
	w.WriteMsg(m)
}

func serve(net, name, secret string) {
	switch name {
	case "":
		server := &dns.Server{Pool: *pool, Addr: ":8053", Net: net, TsigSecret: nil}
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Failed to setup the "+net+" server: %s\n", err.Error())
		}
	default:
		server := &dns.Server{Pool: *pool, Addr: ":8053", Net: net, TsigSecret: map[string]string{name: secret}}
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Failed to setup the "+net+" server: %s\n", err.Error())
		}
	}
}

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	printf = flag.Bool("print", false, "print replies")
	compress = flag.Bool("compress", false, "compress replies")
	pool = flag.Bool("pool", false, "use UDP memory pooling")
	tsig = flag.String("tsig", "", "use MD5 hmac tsig: keyname:base64")
	var name, secret string
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()
	if *tsig != "" {
		a := strings.SplitN(*tsig, ":", 2)
		name, secret = dns.Fqdn(a[0]), a[1] // fqdn the name, which everybody forgets...
	}
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	dns.HandleFunc(".", handleQ)
	go serve("tcp", name, secret)
	go serve("udp", name, secret)
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
forever:
	for {
		select {
		case s := <-sig:
			fmt.Printf("Signal (%d) received, stopping\n", s)
			break forever
		}
	}
}

package main

import (
	"fmt"
	"github.com/miekg/dns"
	"testing"
	"time"
)

func inline() (*dns.Client, *dns.Msg) {
	c := new(dns.Client)
	m := new(dns.Msg)
	return c, m
}

func inline_result(r *dns.Msg, rtt time.Duration, err error) {
	fmt.Printf("Request got answer finished in %s \n", rtt)
	if err != nil {
		fmt.Printf("Huston, we have problems, %s \n", err)
	}
	fmt.Printf("Answer is: %s \n", r.String())
	fmt.Printf("TXT is: %s \n", r.Answer[0].String())
}

func TestResponseTXT(t *testing.T) {
	c, m := inline()
	m.SetQuestion(dns.Fqdn("testing^_^"), dns.TypeTXT)
	in, rtt, err := c.Exchange(m, "127.0.0.1:8053")
	inline_result(in, rtt, err)
	if in.Answer[0].String() != ".	0	CLASS0	TYPE0	\\# 2 4f4b" { //seems like OK?
		t.Fail()
	}
}

func TestResponseAAAA(t *testing.T) {
	c, m := inline()
	m.SetQuestion(dns.Fqdn("localhost"), dns.TypeAAAA)
	in, rtt, err := c.Exchange(m, "127.0.0.1:8053")
	inline_result(in, rtt, err)
	if in.Answer[0].String() != ".	0	CLASS0	TYPE0	\\# 15 000000000000000000000000000001" { //seems like OK?
		t.Fail()
	}
}

package main

import (
	"flag"
	"net"
	"strings"
	"time"

	"github.com/miekg/dns"
)

var (
	queryType        = flag.String("t", "a", "query type")
	clientIP         = flag.String("cip", "", "client ip")
	dnsServer        = flag.String("svr", "8.8.8.8", "dns server")
	recursionDesired = flag.Bool("rd", true, "recursion desired")
)

func resolve(server string, domain string, clientIp *string) ([]dns.RR, error) {
	// queryType
	var qtype uint16
	qtype = dns.TypeA

	// dnsServer
	if !strings.HasSuffix(server, ":53") {
		server += ":53"
	}

	domain = dns.Fqdn(domain)

	msg := new(dns.Msg)
	msg.SetQuestion(domain, qtype)
	msg.RecursionDesired = true

	if *clientIP != "" {
		opt := new(dns.OPT)
		opt.Hdr.Name = "."
		opt.Hdr.Rrtype = dns.TypeOPT
		e := new(dns.EDNS0_SUBNET)
		e.Code = dns.EDNS0SUBNET
		e.Family = 1 // ipv4
		e.SourceNetmask = 32
		e.SourceScope = 0
		e.Address = net.ParseIP(*clientIP).To4()
		opt.Option = append(opt.Option, e)
		msg.Extra = []dns.RR{opt}
	}

	client := &dns.Client{
		DialTimeout:  5 * time.Second,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	resp, rtt, err := client.Exchange(msg, server)
	return resp.Answer, err
}

package net_test

import (
	"context"
	"fmt"
	"net"
	"time"
)

var archWiki = "wiki.archlinux.org"

func exampleResolver() *net.Resolver {
	return &net.Resolver{
		PreferGo:     true,
		StrictErrors: false,
		Dial:         nil,
	}
}

func exampleContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second * 5)
}

func ExampleResolver_LookupAddr() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()
	names, err := r.LookupAddr(ctx, "135.181.27.174")
	if err != nil {
		panic(err)
	}
	for _, name := range names {
		fmt.Println(name)
	}

	// Output:
	// wiki.archlinux.org.
}

func ExampleResolver_LookupCNAME() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()
	cname, err := r.LookupCNAME(ctx, archWiki)
	if err != nil {
		panic(err)
	}

	fmt.Println(cname)
	// Output:
	// wiki.archlinux.org.
}

func ExampleResolver_LookupHost() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()
	addrs, err := r.LookupHost(ctx, archWiki)
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}

	// Output:
	// 135.181.27.174
	// 2a01:4f9:c010:a4eb::1
}

func ExampleResolver_LookupIP() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()

	helper := func(network string) {
		fmt.Printf("network is %s\n", network)
		ips, err := r.LookupIP(ctx, network, archWiki)
		if err != nil {
			panic(err)
		}
		for _, ip := range ips {
			fmt.Println(ip)
		}
	}

	helper("ip")
	helper("ip4")
	helper("ip6")

	// Output:
	// network is ip
	// 135.181.27.174
	// 2a01:4f9:c010:a4eb::1
	// network is ip4
	// 135.181.27.174
	// network is ip6
	// 2a01:4f9:c010:a4eb::1
}

func ExampleResolver_LookupIPAddr() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()
	ipAddrs, err := r.LookupIPAddr(ctx, archWiki)
	if err != nil {
		panic(err)
	}
	for _, ipAddr := range ipAddrs {
		fmt.Println(&ipAddr)
	}

	// Output:
	// 135.181.27.174
	// 2a01:4f9:c010:a4eb::1
}

func ExampleResolver_LookupMX() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()
	mxes, err := r.LookupMX(ctx, "gmail.com") // qq.com 163.com
	if err != nil {
		panic(err)
	}
	for _, mx := range mxes {
		fmt.Printf("%+v\n", *mx)
	}

	// Output:
	// {Host:gmail-smtp-in.l.google.com. Pref:5}
	// {Host:alt1.gmail-smtp-in.l.google.com. Pref:10}
	// {Host:alt2.gmail-smtp-in.l.google.com. Pref:20}
	// {Host:alt3.gmail-smtp-in.l.google.com. Pref:30}
	// {Host:alt4.gmail-smtp-in.l.google.com. Pref:40}
}

func ExampleResolver_LookupNS() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()
	nses, err := r.LookupNS(ctx, "archlinux.org")
	if err != nil {
		panic(err)
	}
	for _, ns := range nses {
		fmt.Printf("%+v\n", *ns)
	}

	// Output:
	// {Host:helium.ns.hetzner.de.}
	// {Host:oxygen.ns.hetzner.com.}
	// {Host:hydrogen.ns.hetzner.com.}
}

func ExampleResolver_LookupNetIP() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()

	helper := func(network string) {
		fmt.Printf("network is %s\n", network)
		ipAddrs, err := r.LookupNetIP(ctx, network, archWiki)
		if err != nil {
			panic(err)
		}
		for _, ipAddr := range ipAddrs {
			fmt.Println(ipAddr)
		}
	}

	helper("ip")
	helper("ip4")
	helper("ip6")

	// Output:
	// network is ip
	// 135.181.27.174
	// 2a01:4f9:c010:a4eb::1
	// network is ip4
	// 135.181.27.174
	// network is ip6
	// 2a01:4f9:c010:a4eb::1
}

func ExampleResolver_LookupPort() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()

	helper := func(network, service string) {
		fmt.Printf("network: %s, service: %s, ", network, service)
		port, err := r.LookupPort(ctx, network, service)
		if err != nil {
			panic(err)
		}
		fmt.Printf("port: %d\n", port)
	}

	helper("tcp", "http")
	helper("udp", "ftp")
	helper("tcp", "hello")

	// Output:
	// network: tcp, service: http, port: 80
	// network: udp, service: ftp, port: 21
	// network: tcp, service: hello, port: 1789
}

// https://www.cloudflare.com/learning/dns/dns-records/dns-srv-record/
func ExampleResolver_LookupSRV() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()

	helper := func(service, proto, name string) {
		fmt.Printf("service: %s, proto: %s, name: %s ==>\n", service, proto, name)
		cname, addrs, err := r.LookupSRV(ctx, service, proto, name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("cname: %s\n", cname)
		for _, addr := range addrs {
			fmt.Printf("%+v\n", addr)
		}
		fmt.Println()
	}

	helper("todo", "todo", "todo")

	// Output:
	// todo
}

// https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/
func ExampleResolver_LookupTXT() {
	ctx, cancel := exampleContext()
	defer cancel()

	r := exampleResolver()

	helper := func(name string) {
		fmt.Printf("%s ==>\n", name)
		txts, err := r.LookupTXT(ctx, name)
		if err != nil {
			panic(err)
		}
		for _, txt := range txts {
			fmt.Println(txt)
		}
		fmt.Println()
	}

	helper("gmail.com")
	helper("163.com")
	helper("qq.com")

	// Output:
	// gmail.com ==>
	// globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=
	// v=spf1 redirect=_spf.google.com
	//
	// 163.com ==>
	// facebook-domain-verification=kqgnezlldheaauy9huiesb3j2emhh3
	// 57c23e6c1ed24f219803362dadf8dea3
	// qdx50vkxg6qpn3n1k6n1tg2syg5wp96y
	// v=spf1 include:spf.mail.163.com -all
	// google-site-verification=hRXfNWRtd9HKlh-ZBOuUgGrxBJh526R8Uygp0jEZ9wY
	//
	// qq.com ==>
	// v=spf1 include:spf.mail.qq.com -all
}
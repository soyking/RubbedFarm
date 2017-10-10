package main

import (
	"github.com/miekg/dns"
	"log"
	"flag"
)

var (
	port     string
	upstream string
)

func init() {
	flag.StringVar(&port, "p", "5400", "listen port")
	flag.StringVar(&upstream, "u", "127.0.0.1:5300", "upstream port")
}

func main() {
	flag.Parse()
	println("Listen on: " + port)
	println("Upstream: " + upstream)

	server := &dns.Server{Addr: "0.0.0.0:" + port, Net: "udp"}
	dns.HandleFunc(".", handleRequest)
	log.Fatal(server.ListenAndServe())
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	c := &dns.Client{
		Net: "tcp",
	}
	m, _, err := c.Exchange(r, upstream)
	if err != nil {
		println("Extrange Error: ", err.Error())
		w.WriteMsg(r)
	} else {
		w.WriteMsg(m)
	}
}

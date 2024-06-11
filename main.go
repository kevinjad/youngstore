package main

import (
	"log"

	"github.com/kevinjad/youngstore/p2p"
)

func main() {
	tcpTransport := p2p.NewTcpTransport(":4000")
	err := tcpTransport.ListenAndAccept()
	if err != nil {
		log.Fatal(err)
	}
	select {}
}

package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TcpPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTcpPeer(conn net.Conn, outbound bool) *TcpPeer {
	return &TcpPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TcpTransport struct {
	listenerAddress string
	listener        net.Listener
	shakeHands      HandShaker
	decoder         Decoder
	mu              sync.RWMutex
	peers           map[net.Addr]TcpPeer
}

func NewTcpTransport(listenerAddress string) *TcpTransport {
	return &TcpTransport{
		listenerAddress: listenerAddress,
		shakeHands:      NOPHandShake,
		decoder:         NOPDecoder{},
	}
}

func (t *TcpTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenerAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TcpTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("error occured in tcp accept %s\n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TcpTransport) handleConn(conn net.Conn) {
	peer := NewTcpPeer(conn, false)
	fmt.Printf("new incoming connection %+v\n", peer)

	if err := t.shakeHands(peer); err != nil {
		fmt.Printf("error occured in handshake for peer %+v\n", peer)
		return
	}

	for {
		msg := make([]byte, 1028)
		t.decoder.Decode(conn, msg)

		fmt.Println(string(msg))
	}
}

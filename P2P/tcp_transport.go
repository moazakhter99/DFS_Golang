package p2p

import (
	"fmt"
	"net"
)

// struct is the remote node over a TCP established connection
type TCPeer struct {
	conn net.Conn

	// If we dial (Outbound) -> true
	// If we Accept (Inbound) -> false
	outbound bool
	//* Could be state with string outbound or inbound
	//* ConnState string

}

func NewTCPeer(conn net.Conn, outbound bool) *TCPeer {
	return &TCPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Implements peer interface
func (peer *TCPeer) Close() error {
	return peer.conn.Close()
}

func OnPeer(Peer) error {
	fmt.Println("Doing some logic with the peer outside of TCPTransport")
	return nil
}

type TCPTranspotOpts struct {
	ListenAddress string
	ShakeHands    HandshakeFunc
	Decoder       Decoder

	OnPeer func(Peer) error
}

type TCPTransport struct {
	TCPTranspotOpts
	Listner net.Listener
	RPCchan chan RPC
}

func NewTCPTransport(opts TCPTranspotOpts) *TCPTransport {
	return &TCPTransport{
		TCPTranspotOpts: opts,
		RPCchan: make(chan RPC),
	}
}

// Implements Transport Interface whicj will return read only channel
// Reading incomming messages from another peer in the network
func (t *TCPTransport) Consume() (<-chan RPC) {
	return t.RPCchan
}


func (t *TCPTransport) ListenAndAccept() (err error) {
	t.Listner, err = net.Listen("tcp", t.ListenAddress)

	if err != nil {
		return
	}

	go t.startAcceptLoop()

	return

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		// First we connect
		conn, err := t.Listner.Accept()
		if err != nil {
			fmt.Printf("TCP Accept Error %s\n", err)
		}

		go t.handdleConn(conn)
	}
}

func (t *TCPTransport) handdleConn(conn net.Conn) {
	var err error
	defer func ()  {
		fmt.Println("Dropping Connection : ", err)
		conn.Close()	
	}()

	// second we Peer
	peer := NewTCPeer(conn, true)

	err = t.ShakeHands(peer)
	if err != nil {
		return
	}

	// Third we check for OnPeer
	if t.OnPeer != nil {
		err = t.OnPeer(peer)
		if err != nil {
			return
		}
	}

	// Lastly we reed
	// Read Loop
	rpc := RPC{}
	for {
		err := t.Decoder.Decode(conn, &rpc)
		if err != nil {
			fmt.Println("Decode Error ; ", err)
			return
		}

		rpc.From = conn.RemoteAddr()
		t.RPCchan <- rpc
	
	}

}

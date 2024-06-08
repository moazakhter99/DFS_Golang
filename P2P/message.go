package p2p

import "net"

// struct holds data to be sent between nodes
type RPC struct {
	From    net.Addr
	Paylaod []byte
}

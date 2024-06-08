package p2p

// Interface that represents the remote node
type Peer interface {
	Close() error

}

// Interface that handles the communication between the nodes in the network.
// This can be of the form (TCP, UDB, websockets, ...)
type Transport interface {

	ListenAndAccept() (error)
	Consume() (<-chan RPC)

}
package p2p

type HandshakeFunc func(interface{}) error

func NOPHandshake(interface{}) error {return nil}
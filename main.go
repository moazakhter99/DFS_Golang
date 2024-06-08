package main

import (
	"fmt"
	"log"

	p2p "github.com/moazakhter99/DFS_Golang/P2P"
)


func main() {
	fmt.Println("Hello world")

	tcpOpts := p2p.TCPTranspotOpts{
		ListenAddress: ":3000",
		ShakeHands: p2p.NOPHandshake,
		Decoder: p2p.DefaultDecoder{},

		OnPeer: p2p.OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func ()  {
		for {
		msg := <- tr.Consume()
		
		fmt.Println("Message : ", msg)
		}
		
	}()

	err := tr.ListenAndAccept()

	if err != nil {
		log.Fatal(err)
	}


	select {}
}
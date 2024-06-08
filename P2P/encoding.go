package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(reader io.Reader, rpc *RPC) (err error) {
	err = gob.NewDecoder(reader).Decode(rpc)

	return

}

type DefaultDecoder struct {
}

func (def DefaultDecoder) Decode(reader io.Reader, rpc *RPC) (err error) {
	buf := make([]byte, 1028)
	n, err := reader.Read(buf)
	if err != nil {
		return
	}

	rpc.Paylaod = buf[:n]

	return
}

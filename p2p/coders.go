package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(r io.Reader, buf []byte) error
}
type DefaultDecoder struct{}

func (d DefaultDecoder) Decode(r io.Reader, buf []byte) error {
	return gob.NewDecoder(r).Decode(&buf)
}

type NOPDecoder struct{}

func (d *NOPDecoder) Decode(r io.Reader, buf []byte) error {
	_, err := r.Read(buf)
	return err
}

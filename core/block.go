package core

import (
	"encoding/binary"
	"goblock/types"
	"io"
)

type Header struct {
	Version   uint32
	Height    uint32
	Timestamp uint64
	Nonce     uint64
	PrevBlock types.Hash
}

func (h *Header) EncodeBinary(w io.Writer) error {
	for _, v := range []interface{}{
		&h.Version,
		&h.Height,
		&h.Timestamp,
		&h.Nonce,
		&h.PrevBlock,
	} {
		if err := binary.Write(w, binary.LittleEndian, v); err != nil {
			return err
		}
	}

	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {
	for _, v := range []interface{}{
		&h.Version,
		&h.Height,
		&h.Timestamp,
		&h.Nonce,
		&h.PrevBlock,
	} {
		if err := binary.Read(r, binary.LittleEndian, v); err != nil {
			return err
		}
	}

	return nil
}

type Block struct {
	Header
	Transactions []Transaction
}

package core

import "goblock/types"

type Header struct {
	Version   uint32
	Timestamp uint64
	Height    uint32
	Nonce     uint64
	PrevBlock types.Hash
}

type Block struct {
	Header
	Transactions []Transaction
}

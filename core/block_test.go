package core

import (
	"bytes"
	"goblock/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeaderEncodeAndDecode(t *testing.T) {

	header := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     908692,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, header.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))

	assert.Equal(t, header, hDecode)
}
func TestBlockEncodeAndDecode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     908692,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))

	assert.Equal(t, b, bDecode)
}

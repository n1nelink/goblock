package core

import (
	"bytes"
	"goblock/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeaderEncode(t *testing.T) {

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
	assert.Nil(t, hDecode.EncodeBinary(buf))

	assert.Equal(t, header, hDecode)
}

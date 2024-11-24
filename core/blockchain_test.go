package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	bc := NewBlockchain()
	assert.NotNil(t, bc.validator)
}

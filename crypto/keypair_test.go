package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAddressFromPrivateKeyAndPublicKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	address := pubKey.Address()
	fmt.Println(address)
}

func TestKeypairSignVerifyValid(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	msg := []byte("Hello, I am message")
	signature, err := privateKey.Sign(msg)

	assert.Nil(t, err)
	assert.True(t, signature.Verify(publicKey, msg))
}

func TestKeypairSignVerifyInvaid(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	otherPrivateKey := GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()

	msg := []byte("Hello, I am right message")
	otherMsg := []byte("Hello, I am not right message")
	signature, err := privateKey.Sign(msg)

	assert.Nil(t, err)
	assert.False(t, signature.Verify(otherPublicKey, msg))
	assert.False(t, signature.Verify(publicKey, otherMsg))
}

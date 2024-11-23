package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"goblock/types"
	"math/big"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{key}
}

func (priv PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &priv.key.PublicKey,
	}
}

func (priv PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, priv.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{r, s}, nil
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (pub PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(pub.key, pub.key.X, pub.key.Y)
}

func (pub PublicKey) Address() types.Address {
	h := sha256.Sum256(pub.ToSlice())
	return types.AddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	r *big.Int
	s *big.Int
}

func (sig Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(pubKey.key, data, sig.r, sig.s)
}

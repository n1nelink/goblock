package core

type Blockchain struct {
	headers   []*Header
	store     Storage
	validator Validator
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{
		headers: []*Header{},
	}
	blockchain.validator = NewBlockValidator(blockchain)

	return blockchain
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	return nil
}

func (bc *Blockchain) Height() uint32 {
	// [0,1,2,3] => (len = 4)
	// [0,1,2,3] => (height = len - 1 = 3)
	return uint32(len(bc.headers) - 1)
}

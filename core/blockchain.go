package core

type Blockchain struct {
	headers   []*Header
	store     Storage
	validator Validator
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	blockchain := &Blockchain{
		headers: []*Header{},
	}
	blockchain.validator = NewBlockValidator(blockchain)

	err := blockchain.addBlockWithoutValidation(genesis)

	return blockchain, err
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	return nil
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height < bc.Height()
}

func (bc *Blockchain) Height() uint32 {
	// [0,1,2,3] => (len = 4)
	// [0,1,2,3] => (height = len - 1 = 3)
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)

	return bc.store.Put(b)
}

func (bc *Blockchain) addGenesisBlock(b *Block) {

}

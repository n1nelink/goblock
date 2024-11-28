package core

type Storage interface {
	Put(b *Block) error
}

type MemoryStore struct {
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Put(b *Block) error {
	return nil
}

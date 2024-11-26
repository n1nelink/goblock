package core

type Storage interface {
	Put(b *Block) error
}

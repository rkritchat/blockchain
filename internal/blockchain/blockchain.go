package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block)DeriveHash(){
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block{
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string){
	prevBlock := chain.Blocks[(len(chain.Blocks)-1)] //hooking previous block from arrays
	n := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, n)
}

func Genesis() *Block{
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain()*Blockchain{
	return &Blockchain{[]*Block{Genesis()}}
}
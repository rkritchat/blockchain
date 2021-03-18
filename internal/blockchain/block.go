package blockchain

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}


func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[(len(chain.Blocks) - 1)] //hooking previous block from arrays
	n := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, n)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

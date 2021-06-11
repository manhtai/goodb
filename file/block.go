package file

type Block struct {
	filename string
	number   int
}

func NewBlock(filename string, blockNumber int) *Block {
	return &Block{filename: filename, number: blockNumber}
}
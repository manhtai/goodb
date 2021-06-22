package file

type Block struct {
	filename string
	number   int
}

func NewBlock(filename string, blockNumber int) *Block {
	return &Block{filename: filename, number: blockNumber}
}

func IsBlocksEq(b1 *Block, b2 *Block) bool {
	return b1.number == b2.number && b1.filename == b2.filename
}

func (b *Block) Filename() string {
	return b.filename
}

func (b *Block) Number() int {
	return b.number
}

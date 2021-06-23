package file

type Block struct {
	filename string
	number   int
}

func NewBlock(filename string, blockNumber int) Block {
	return Block{
		filename: filename,
		number:   blockNumber,
	}
}

func (b Block) Filename() string {
	return b.filename
}

func (b Block) Number() int {
	return b.number
}

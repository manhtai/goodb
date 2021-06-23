package file

import "goodb/constant"

type Page struct {
	buffer []byte
}

func NewPage(blockSize int) *Page {
	return &Page{
		buffer: make([]byte, blockSize),
	}
}

func NewPageFromBytes(bytes []byte) *Page {
	return &Page{
		buffer: bytes,
	}
}

func (p *Page) GetInt(offset int) int {
	b := p.buffer[offset]
	return int(b)
}

func (p *Page) SetInt(offset int, n int) {
	p.buffer[offset] = byte(n)
}

func (p *Page) GetBytes(offset int) []byte {
	length := p.GetInt(offset)
	bytes := p.buffer[offset+constant.INT_SIZE : offset+length+constant.INT_SIZE]
	return bytes
}

func (p *Page) SetBytes(offset int, bytes []byte) {
	p.SetInt(offset, len(bytes))
	for i := 0; i < len(bytes); i++ {
		p.buffer[offset+i+constant.INT_SIZE] = bytes[i]
	}
}

func (p *Page) GetString(offset int) string {
	bytes := p.GetBytes(offset)
	return string(bytes)
}

func (p *Page) SetString(offset int, s string) {
	bytes := []byte(s)
	p.SetBytes(offset, bytes)
}

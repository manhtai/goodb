package file

type Page struct {
	buffer []byte
}

func (p *Page) GetInt(offset int) int {
	b := p.buffer[offset]
	return int(b)
}

func (p *Page) SetInt(offset int, n int) {
	p.buffer[offset] = byte(n)
}

func (p *Page) SetBytes(offset int, bytes []byte) {
	for off := 0; off < len(bytes); off++ {
		p.buffer[offset+off] = bytes[off]
	}
}

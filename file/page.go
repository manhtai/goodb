package file

type Page struct {
	buffer []byte
}

func (p *Page) getInt(offset int) int {
	b := p.buffer[offset]
	return int(b)
}

func (p *Page) setInt(offset int, n int) {
	p.buffer[offset] = byte(n)
}

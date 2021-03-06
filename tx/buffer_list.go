package tx

import (
	"goodb/buffer"
	"goodb/file"
)

type BufferList struct {
	buffers   map[file.Block]*buffer.Buffer
	pins      []file.Block
	bufferMgr *buffer.BufferManager
}

func NewBufferList(bufferMgr *buffer.BufferManager) *BufferList {
	return &BufferList{
		bufferMgr: bufferMgr,
		buffers:   make(map[file.Block]*buffer.Buffer),
		pins:      make([]file.Block, 1),
	}
}

func (bl *BufferList) getBuffer(block file.Block) *buffer.Buffer {
	return bl.buffers[block]
}

func (bl *BufferList) pin(block file.Block) {
	buff := bl.bufferMgr.Pin(block)
	bl.buffers[block] = buff
	bl.pins = append(bl.pins, block)
}

func (bl *BufferList) unpin(block file.Block) {
	buff := bl.buffers[block]
	bl.bufferMgr.Unpin(buff)

	pinIndex := bl.findPinIndex(block)
	if pinIndex > -1 {
		bl.pins = append(bl.pins[:pinIndex], bl.pins[pinIndex+1:]...)
	}

	if bl.findPinIndex(block) == -1 {
		delete(bl.buffers, block)
	}
}

func (bl *BufferList) findPinIndex(block file.Block) int {
	for i, b := range bl.pins {
		if b == block {
			return i
		}
	}
	return -1
}

func (bl *BufferList) unpinAll() {
	for _, blk := range bl.pins {
		if blk.Filename() != "" {
			buff := bl.buffers[blk]
			bl.bufferMgr.Unpin(buff)
		}
	}
	bl.buffers = make(map[file.Block]*buffer.Buffer)
	bl.pins = bl.pins[:0]
}

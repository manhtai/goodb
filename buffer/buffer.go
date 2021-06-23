package buffer

import (
	"goodb/file"
	"goodb/log"
)

type Buffer struct {
	fileMgr  *file.FileManager
	logMgr   *log.LogManager
	contents *file.Page
	block    file.Block
	pins     int
	txNum    int
	lsn      int
}

func NewBuffer(fileMgr *file.FileManager, logMgr *log.LogManager) *Buffer {
	return &Buffer{
		fileMgr:  fileMgr,
		logMgr:   logMgr,
		contents: file.NewPage(fileMgr.BlockSize()),
	}
}

func (buffer *Buffer) Contents() *file.Page {
	return buffer.contents
}

func (buffer *Buffer) Block() file.Block {
	return buffer.block
}

func (buffer *Buffer) SetModified(txNum int, lsn int) {
	buffer.txNum = txNum
	if lsn >= 0 {
		buffer.lsn = lsn
	}
}

func (buffer *Buffer) IsPinned() bool {
	return buffer.pins > 0
}

func (buffer *Buffer) assignToBlock(block file.Block) {
	buffer.flush()
	buffer.block = block
	buffer.fileMgr.Read(block, buffer.contents)
	buffer.pins = 0
}

func (buffer *Buffer) flush() {
	if buffer.txNum >= 0 && buffer.block.Filename() != "" {
		buffer.logMgr.Flush(buffer.lsn)
		buffer.fileMgr.Write(buffer.block, buffer.contents)
		buffer.txNum = -1
	}
}

func (buffer *Buffer) pin() {
	buffer.pins += 1
}

func (buffer *Buffer) unpin() {
	buffer.pins -= 1
}

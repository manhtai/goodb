package buffer

import (
	"goodb/file"
	"goodb/log"
	"time"
)

const WAIT_TIME = 10 * time.Second

type BufferManager struct {
	bufferPool   []*Buffer
	numAvailable int
}

func NewBufferManager(fileMgr *file.FileManager, logMgr *log.LogManager, numBuffers int) *BufferManager {
	bufferPool := make([]*Buffer, numBuffers)
	for i := 0; i < numBuffers; i++ {
		bufferPool[i] = NewBuffer(fileMgr, logMgr)
	}
	return &BufferManager{
		bufferPool:   bufferPool,
		numAvailable: numBuffers,
	}
}

func (bufferMgr *BufferManager) GetNumAvailable() int {
	return bufferMgr.numAvailable
}

func (bufferMgr *BufferManager) FlushAll(txNum int) {
	for _, buffer := range bufferMgr.bufferPool {
		if buffer.txNum == txNum {
			buffer.flush()
		}
	}
}

func (bufferMgr *BufferManager) Unpin(buffer *Buffer) {
	buffer.unpin()
	if !buffer.IsPinned() {
		bufferMgr.numAvailable += 1
	}
}

func (bufferMgr *BufferManager) Pin(block file.Block) *Buffer {
	buffer := bufferMgr.tryToPin(block)
	now := time.Now()
	for buffer == nil && time.Since(now) < WAIT_TIME {
		buffer = bufferMgr.tryToPin(block)
		time.Sleep(time.Second)
	}
	if buffer == nil {
		panic("can't pin block")
	}
	return buffer
}

func (bufferMgr *BufferManager) tryToPin(block file.Block) *Buffer {
	buffer := bufferMgr.findExistingBuffer(block)
	if buffer == nil {
		buffer = bufferMgr.chooseUnpinnedBuffer()
		if buffer == nil {
			return nil
		}
		buffer.assignToBlock(block)
	}
	if !buffer.IsPinned() {
		bufferMgr.numAvailable -= 1
	}
	buffer.pin()
	return buffer
}

func (bufferMgr *BufferManager) findExistingBuffer(block file.Block) *Buffer {
	for _, buffer := range bufferMgr.bufferPool {
		if buffer.block == block {
			return buffer
		}
	}
	return nil
}

func (bufferMgr *BufferManager) chooseUnpinnedBuffer() *Buffer {
	for _, buffer := range bufferMgr.bufferPool {
		if !buffer.IsPinned() {
			return buffer
		}
	}
	return nil
}

package log

import (
	"goodb/file"
)

const INT_SIZE = 4

type LogManager struct {
	fileMgr *file.FileManager
	logFile string
	logPage *file.Page
	currentBlock *file.Block
	latestLSN int
	lastSavedLSN int
}

func (logMgr *LogManager) Append(logRecord []byte) {
	boundary := logMgr.logPage.GetInt(0)
	recordSize := len(logRecord)
	bytesNeeded := recordSize + INT_SIZE

	if boundary - bytesNeeded < INT_SIZE {
		logMgr.flush()
		logMgr.currentBlock = logMgr.appendNewBlock()
		boundary = logMgr.logPage.GetInt(0)
	}

	recordPosition := boundary - bytesNeeded
	logMgr.logPage.SetBytes(recordPosition, logRecord)
	logMgr.logPage.SetInt(0, recordPosition)
	logMgr.latestLSN += 1
}

func (logMgr *LogManager) appendNewBlock() *file.Block {
	block := logMgr.fileMgr.Append(logMgr.logFile)
	logMgr.logPage.SetInt(0, logMgr.fileMgr.GetBlockSize())
	logMgr.fileMgr.Write(block, logMgr.logPage)
	return block
}

func (logMgr *LogManager) Flush(lsn int) {
	if lsn >= logMgr.lastSavedLSN {
		logMgr.flush()
	}
}

func (logMgr *LogManager) flush() {
	logMgr.fileMgr.Write(logMgr.currentBlock, logMgr.logPage)
	logMgr.lastSavedLSN = logMgr.latestLSN
}

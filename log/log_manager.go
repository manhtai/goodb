package log

import (
	"goodb/constant"
	"goodb/file"
)

type LogManager struct {
	fileMgr      *file.FileManager
	logFile      string
	logPage      *file.Page
	currentBlock file.Block
	latestLSN    int
	lastSavedLSN int
}

func NewLogManager(fileMgr *file.FileManager, logFile string) *LogManager {
	bytes := make([]byte, fileMgr.BlockSize())
	logPage := file.NewPageFromBytes(bytes)
	logSize := fileMgr.Length(logFile)

	logMgr := &LogManager{
		fileMgr: fileMgr,
		logFile: logFile,
		logPage: logPage,
	}

	var currentBlock file.Block
	if logSize == 0 {
		currentBlock = logMgr.appendNewBlock()
	} else {
		currentBlock = file.NewBlock(logFile, logSize-1)
		fileMgr.Read(currentBlock, logPage)
	}

	logMgr.currentBlock = currentBlock
	return logMgr
}

func (logMgr *LogManager) Append(logRecord []byte) int {
	boundary := logMgr.logPage.GetInt(0)
	recordSize := len(logRecord)
	bytesNeeded := recordSize + constant.INT_SIZE

	if boundary-bytesNeeded < constant.INT_SIZE {
		logMgr.flush()
		logMgr.currentBlock = logMgr.appendNewBlock()
		boundary = logMgr.logPage.GetInt(0)
	}

	recordPosition := boundary - bytesNeeded
	logMgr.logPage.SetBytes(recordPosition, logRecord)
	logMgr.logPage.SetInt(0, recordPosition)
	logMgr.latestLSN += 1
	return logMgr.latestLSN
}

func (logMgr *LogManager) appendNewBlock() file.Block {
	block := logMgr.fileMgr.Append(logMgr.logFile)
	logMgr.logPage.SetInt(0, logMgr.fileMgr.BlockSize())
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

func MaxLength(strLen int) int {
	return constant.INT_SIZE + strLen*constant.INT_SIZE
}

package file

import (
	"io"
	"os"
	"path"
)

type FileManager struct {
	dbDirectory string
	blockSize int
	isNew bool
	openFiles map[string]*os.File
}

func (fileMgr *FileManager) read(blockId *BlockId, page *Page) {
	file := fileMgr.readFile(blockId.filename)
	_, err := file.Seek(int64(blockId.number*fileMgr.blockSize), 0)
	if err != nil {
		panic(err)
	}
	_, err = io.ReadAtLeast(file, page.buffer, 0)
	if err != nil {
		panic(err)
	}
}

func (fileMgr *FileManager) write(blockId *BlockId, page *Page) {
	file := fileMgr.readFile(blockId.filename)
	_, err := file.WriteAt(page.buffer, int64(blockId.number*fileMgr.blockSize))
	if err != nil {
		panic(err)
	}
}

func (fileMgr *FileManager) readFile(filename string) *os.File {
	file, ok := fileMgr.openFiles[filename]
	if ok {
		return file
	}
	filePath := path.Join(fileMgr.dbDirectory, filename)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModeExclusive)
	if err != nil {
		panic(err)
	}
	fileMgr.openFiles[filename] = file
	return file
}
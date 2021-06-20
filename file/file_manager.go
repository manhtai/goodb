package file

import (
	"io"
	"os"
	"path"
)

type FileManager struct {
	dbDirectory string
	blockSize   int
	isNew       bool
	openFiles   map[string]*os.File
}

func NewFileManager(dbDir string, blockSize int) *FileManager {
	isNew := false
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		isNew = true
		err := os.Mkdir(dbDir, os.ModeDir)
		if err != nil {
			panic("create db directory failure")
		}
	}

	return &FileManager{
		dbDirectory: dbDir,
		blockSize:   blockSize,
		isNew:       isNew,
		openFiles:   make(map[string]*os.File),
	}
}

func (fileMgr *FileManager) BlockSize() int {
	return fileMgr.blockSize
}

func (fileMgr *FileManager) Read(block *Block, page *Page) {
	file := fileMgr.getFile(block.filename)
	_, err := file.Seek(int64(block.number*fileMgr.blockSize), 0)
	if err != nil {
		panic(err)
	}
	_, err = io.ReadAtLeast(file, page.buffer, 0)
	if err != nil {
		panic(err)
	}
}

func (fileMgr *FileManager) Write(block *Block, page *Page) {
	file := fileMgr.getFile(block.filename)
	_, err := file.WriteAt(page.buffer, int64(block.number*fileMgr.blockSize))
	if err != nil {
		panic(err)
	}
}

func (fileMgr *FileManager) Append(filename string) *Block {
	newBlockNumber := len(filename)
	block := &Block{filename: filename, number: newBlockNumber}
	b := make([]byte, fileMgr.blockSize)
	file := fileMgr.getFile(block.filename)
	_, err := file.WriteAt(b, int64(block.number*fileMgr.blockSize))
	if err != nil {
		panic(err)
	}
	return block
}

func (fileMgr *FileManager) getFile(filename string) *os.File {
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

func (fileMgr *FileManager) Length(filename string) int {
	file := fileMgr.getFile(filename)
	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}
	return int(fi.Size())
}

func (fileMgr *FileManager) IsNew() bool {
	return fileMgr.isNew
}

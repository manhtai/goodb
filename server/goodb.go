package server

import (
	"fmt"
	"goodb/buffer"
	"goodb/file"
	"goodb/log"
	"goodb/metadata"
	"goodb/plan"
	"goodb/plan/basic"
	"goodb/tx"
)

const (
	BLOCK_SIZE  = 400
	BUFFER_SIZE = 8
	LOG_FILE    = "gooDb.log"
)

type GooDb struct {
	fileMgr     *file.FileManager
	bufferMgr   *buffer.BufferManager
	logMgr      *log.LogManager
	metadataMgr *metadata.MetadataManager
	planner     *plan.Planner
}

func NewGooDb(dirName string) *GooDb {
	fileMgr := file.NewFileManager(dirName, BLOCK_SIZE)
	logMgr := log.NewLogManager(fileMgr, LOG_FILE)
	bufferMgr := buffer.NewBufferManager(fileMgr, logMgr, BUFFER_SIZE)
	gooDb := &GooDb{
		fileMgr:   fileMgr,
		bufferMgr: bufferMgr,
		logMgr:    logMgr,
	}

	transaction := tx.NewTransaction(fileMgr, logMgr, bufferMgr)

	isNew := fileMgr.IsNew()
	if isNew {
		fmt.Println("Create new database")
	} else {
		fmt.Println("Recover old database")
		transaction.Recover()
	}
	metadataMgr := metadata.NewMetadataManager(isNew, transaction)
	queryPlanner := basic.NewBasicQueryPlanner(metadataMgr)
	updatePlanner := basic.NewBasicUpdatePlanner(metadataMgr)
	planner := plan.NewPlanner(queryPlanner, updatePlanner)

	gooDb.metadataMgr = metadataMgr
	gooDb.planner = planner

	transaction.Commit()
	return gooDb
}

func (db *GooDb) NewTx() *tx.Transaction {
	return tx.NewTransaction(db.fileMgr, db.logMgr, db.bufferMgr)
}

func (db *GooDb) Planner() *plan.Planner {
	return db.planner
}

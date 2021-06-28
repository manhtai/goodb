package server

import (
	"fmt"
	"goodb/buffer"
	"goodb/file"
	"goodb/log"
	"goodb/metadata"
	"goodb/plan"
	"goodb/plan/basic"
	"goodb/plan/opt"
	"goodb/tx"
)

const (
	BLOCK_SIZE  = 400
	BUFFER_SIZE = 8
	LOG_FILE    = "goodb.log"
)

type GooDb struct {
	fileMgr     *file.FileManager
	bufferMgr   *buffer.BufferManager
	logMgr      *log.LogManager
	metadataMgr *metadata.MetadataManager
	planner     *plan.Planner
}

func NewGooDbBasic(dirName string) *GooDb {
	return NewGooDbWithPlanner(dirName, basic.NewBasicQueryPlanner, basic.NewBasicUpdatePlanner)
}

func NewGooDbOpt(dirName string) *GooDb {
	return NewGooDbWithPlanner(dirName, opt.NewOptQueryPlanner, opt.NewIndexUpdatePlanner)
}

func NewGooDbWithPlanner(
	dirName string,
	queryPlannerFunc plan.QueryPlannerFunc,
	updatePlannerFunc plan.UpdatePlannerFunc,
) *GooDb {
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
		fmt.Printf("Create new database: %s\n", dirName)
	} else {
		fmt.Printf("Recover old database: %s\n", dirName)
		transaction.Recover()
	}
	metadataMgr := metadata.NewMetadataManager(isNew, transaction)
	queryPlanner := queryPlannerFunc(metadataMgr)
	updatePlanner := updatePlannerFunc(metadataMgr)
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

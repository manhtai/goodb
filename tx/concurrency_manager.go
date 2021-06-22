package tx

import "goodb/file"

type ConcurrencyManager struct {
	lockTable *LockTable
	locks     map[*file.Block]string
}

func (conMgr *ConcurrencyManager) SLock(block *file.Block) {
	if _, ok := conMgr.locks[block]; !ok {
		conMgr.lockTable.SLock(block)
		conMgr.locks[block] = "S"
	}
}

func (conMgr *ConcurrencyManager) XLock(block *file.Block) {
	if !conMgr.hasXLock(block) {
		conMgr.SLock(block)
		conMgr.lockTable.XLock(block)
		conMgr.locks[block] = "X"
	}
}

func (conMgr *ConcurrencyManager) Release() {
	for block := range conMgr.locks {
		conMgr.lockTable.unlock(block)
		delete(conMgr.locks, block)
	}
}

func (conMgr *ConcurrencyManager) hasXLock(block *file.Block) bool {
	lockType, ok := conMgr.locks[block]
	return ok && lockType == "X"
}

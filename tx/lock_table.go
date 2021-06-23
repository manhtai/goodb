package tx

import (
	"goodb/file"
	"time"
)

const WAIT_TIME = 10 * time.Second

type LockTable struct {
	locks map[file.Block]int
}

func NewLockTable() *LockTable {
	return &LockTable{
		locks: make(map[file.Block]int),
	}
}

func (lt *LockTable) SLock(block file.Block) {
	now := time.Now()
	for lt.hasXLock(block) && time.Since(now) < WAIT_TIME {
		time.Sleep(time.Second)
	}
	if lt.hasXLock(block) {
		panic("wait for xlock for too long")
	}
	lt.locks[block] = lt.getLockVal(block) + 1
}

func (lt *LockTable) XLock(block file.Block) {
	now := time.Now()
	for lt.hasOtherSLocks(block) && time.Since(now) < WAIT_TIME {
		time.Sleep(time.Second)
	}
	if lt.hasOtherSLocks(block) {
		panic("wait for slock for too long")
	}
	lt.locks[block] = -1
}

func (lt *LockTable) unlock(block file.Block) {
	val := lt.getLockVal(block)
	if val > 1 {
		lt.locks[block] = val - 1
	} else {
		delete(lt.locks, block)
	}
}
func (lt *LockTable) hasXLock(block file.Block) bool {
	return lt.getLockVal(block) < 0
}
func (lt *LockTable) hasOtherSLocks(block file.Block) bool {
	return lt.getLockVal(block) > 1
}
func (lt *LockTable) getLockVal(block file.Block) int {
	val, ok := lt.locks[block]
	if !ok {
		return 0
	}
	return val
}

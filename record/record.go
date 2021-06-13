package record

type Record struct {
	blockNumber int
	slot int
}

func NewRecord(blockNumber int, slot int) *Record {
	return &Record{
		blockNumber: blockNumber,
		slot: slot,
	}
}

func (r *Record) BlockNumber() int {
	return r.blockNumber
}

func (r *Record) Slot() int {
	return r.slot
}

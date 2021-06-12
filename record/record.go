package record

type Record struct {
	blockNumber int
	slot int
}

func (r *Record) BlockNumber() int {
	return r.blockNumber
}

func (r *Record) Slot() int {
	return r.slot
}

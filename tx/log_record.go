package tx

import (
	"goodb/file"
)

type LogRecord interface {
	op() int
	txNumber() int
	undo(tx *Transaction)
}

const (
	CHECKPOINT int = iota
	START
	COMMIT
	ROLLBACK
	SETINT
	SETSTRING
)

func CreateLogRecord(bytes []byte) LogRecord {
	page := file.NewPageFromBytes(bytes)
	switch page.GetInt(0) {
	case CHECKPOINT:
		return &CheckpointRecord{}
	case START:
	case COMMIT:
	case ROLLBACK:
	case SETINT:
	case SETSTRING:
	}
	return nil
}

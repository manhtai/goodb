package recovery

import (
	"goodb/file"
	"goodb/log"
	"goodb/tx"
)

type SetIntRecord struct {
	txNum int
	offset int
	val int
	block *file.Block
}

func NewSetIntRecord(page *file.Page) *SetIntRecord {
	pos := log.INT_SIZE
	txNum := page.GetInt(pos)

	pos += log.INT_SIZE
	filename := page.GetString(pos)

	pos += len(filename)
	blockNum := page.GetInt(pos)
	block := file.NewBlock(filename, blockNum)

	pos += log.INT_SIZE
	offset := page.GetInt(pos)

	pos += log.INT_SIZE
	val := page.GetInt(pos)

	return &SetIntRecord{
		txNum: txNum,
		offset: offset,
		val: val,
		block: block,
	}
}

func (r *SetIntRecord) op() int {
	return SETINT
}

func (r *SetIntRecord) txNumber() int {
	return r.txNum
}

func (r *SetIntRecord) undo(tx *tx.Transaction) {
	tx.Pin(r.block)
	tx.SetInt(r.block, r.offset, r.val, false)
	tx.Unpin(r.block)
}

/*

   int tpos = Integer.BYTES;
   int fpos = tpos + Integer.BYTES;
   int bpos = fpos + Page.maxLength(blk.fileName().length());
   int opos = bpos + Integer.BYTES;
   int vpos = opos + Integer.BYTES;
   byte[] rec = new byte[vpos + Integer.BYTES];
   Page p = new Page(rec);
   p.setInt(0, SETINT);
   p.setInt(tpos, txnum);
   p.setString(fpos, blk.fileName());
   p.setInt(bpos, blk.number());
   p.setInt(opos, offset);
   p.setInt(vpos, val);
   return lm.append(rec);
 */

func WriteSETINToLog(logMgr *log.LogManager, txNum int, block *file.Block, offset int, val int) int {
	txPos := log.INT_SIZE
	filePos := txPos + log.INT_SIZE
	blockPos := filePos + len(block.Filename())
	offsetPos := blockPos + log.INT_SIZE
	valuePos := offsetPos + log.INT_SIZE

	record := make([]byte, valuePos + log.INT_SIZE)
	page := file.NewPageFromBytes(record)
	page.SetInt(0, SETINT)
	page.SetInt(txPos, txNum)
	page.SetString(filePos, block.Filename())
	page.SetInt(blockPos, block.Number())
	page.SetInt(offsetPos, offset)
	page.SetInt(valuePos, val)

	return logMgr.Append(record)
}
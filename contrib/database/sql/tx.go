package sql

import "database/sql/driver"

type raspTx struct {
	driver.Tx
}

func (t raspTx) Commit() error {
	return t.Tx.Commit()
}

func (t raspTx) Rollback() error {
	return t.Tx.Rollback()
}

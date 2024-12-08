package sql

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"sync"
)

type raspDriver struct {
	driver.Driver
}

var (
	// For type assertion
	_ driver.Driver = &raspDriver{}

	regMu sync.Mutex
)

func Register(driverName string) (string, error) {
	db, err := sql.Open(driverName, "")
	if err != nil {
		return "", err
	}

	driver := db.Driver()
	if err = db.Close(); err != nil {
		return "", err
	}

	regMu.Lock()
	defer regMu.Unlock()
	name := fmt.Sprintf("%s-gorasp-%d", driverName, len(sql.Drivers()))
	sql.Register(name, driver)

	return name, nil
}

func (d raspDriver) Open(name string) (driver.Conn, error) {
	return d.Open(name)
}

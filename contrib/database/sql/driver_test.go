package sql_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	raspSql "github.com/mrtc0/gorasp/contrib/database/sql"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Parallel()

	driverName, err := raspSql.Register("sqlite3")
	assert.NoError(t, err)

	_, err = sql.Open(driverName, "file:test.db?cache=shared&mode=memory")
	assert.NoError(t, err)
}

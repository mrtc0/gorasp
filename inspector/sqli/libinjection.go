package sqli

import (
	"fmt"

	"github.com/mrtc0/gorasp/lib"
)

type SQLiInspectArgs struct {
	Value string
}

func Inspect(value SQLiInspectArgs) error {
	var fingerprint string
	isSQLi := lib.LibinjectionSQLiFunc(value.Value, len(value.Value), fingerprint)
	if isSQLi == 1 {
		return fmt.Errorf("SQLi detected")
	}

	return nil
}

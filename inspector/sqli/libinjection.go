package sqli

import (
	"fmt"

	"github.com/mrtc0/gorasp/lib"
)

func IsSQLi(value string) error {
	var fingerprint string
	isSQLi := lib.LibinjectionSQLiFunc(value, len(value), fingerprint)
	if isSQLi == 1 {
		return fmt.Errorf("SQLi detected")
	}

	return nil
}

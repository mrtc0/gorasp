package sqli

import "fmt"

func IsSQLiQuery(query string) error {
	isSQLi, err := IsWhereTautologyFull(query)
	if err != nil {
		return err
	}

	if isSQLi {
		return fmt.Errorf("SQLi detected")
	}

	return nil
}

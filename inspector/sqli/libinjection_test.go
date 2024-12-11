package sqli_test

import (
	"testing"

	"github.com/mrtc0/gorasp/inspector/sqli"
	"github.com/mrtc0/gorasp/lib"
	"github.com/stretchr/testify/assert"
)

func TestIsSQLiPayload(t *testing.T) {
	t.Parallel()

	lib.Load()

	testCases := []struct {
		value       string
		expectError bool
	}{
		{
			value:       "test",
			expectError: false,
		},
		{
			value:       "-1' and 1=1 union/* foo */select load_file('/etc/passwd')--",
			expectError: true,
		},
		{
			value:       "test' # ",
			expectError: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()

			err := sqli.IsSQLiPayload(tc.value)
			if tc.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

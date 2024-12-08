package lib

import (
	_ "embed"
	"os"

	"github.com/ebitengine/purego"
)

//go:embed vendor/libinjection/libinjection.so
var LibinjectionSharedLib []byte

var (
	LibinjectionSQLiFunc func(string, int, string) int
)

func Load() (bool, error) {
	const libinjectionSQLiSym = "libinjection_sqli"

	f, err := dumpLibinjection()
	if err != nil {
		return false, err
	}
	defer f.Close()

	handle, err := purego.Dlopen(f.Name(), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return false, err
	}

	purego.RegisterLibFunc(&LibinjectionSQLiFunc, handle, libinjectionSQLiSym)

	return true, nil
}

// https://github.com/ebitengine/purego/issues/102
func dumpLibinjection() (*os.File, error) {
	f, err := os.CreateTemp("", "libinjection.so")
	if err != nil {
		return nil, err
	}

	if _, err := f.Write(LibinjectionSharedLib); err != nil {
		return f, err
	}

	return f, nil
}

package gowaf

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/ebitengine/purego"
	"github.com/mrtc0/gowaf/emitter"
)

//go:embed lib/vendor/libinjection/libinjection.so
var libinjection []byte

var (
	libinjectionSQLi func(string, int, string) int
)

type waf struct{}

func Start() {
	ok, err := Load()
	if !ok || err != nil {
		panic(err)
	}

	eventEmitter := emitter.GetEventEmitter()
	eventEmitter.On("http", func(params emitter.Params) {
		var fingerprint string
		isSQLi := libinjectionSQLi(params.RequestURI, len(params.RequestURI), fingerprint)
		fmt.Println("isSQLi:", isSQLi)
	})
}

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

	purego.RegisterLibFunc(&libinjectionSQLi, handle, libinjectionSQLiSym)

	return true, nil
}

// https://github.com/ebitengine/purego/issues/102
func dumpLibinjection() (*os.File, error) {
	f, err := os.CreateTemp("", "libinjection.so")
	if err != nil {
		return nil, err
	}

	if _, err := f.Write(libinjection); err != nil {
		return f, err
	}

	return f, nil
}

package gorasp

import (
	_ "embed"

	"github.com/mrtc0/gorasp/lib"
)

func Start() {
	ok, err := Load()
	if !ok || err != nil {
		panic(err)
	}
}

func Load() (bool, error) {
	return lib.Load()
}

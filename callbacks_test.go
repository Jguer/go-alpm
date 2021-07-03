package alpm

import (
	"testing"
)

type Cnt struct {
	cnt int
}

func TestCallbacks(t *testing.T) {
	h, _ := Initialize("/", "/var/lib/pacman")
	cnt := &Cnt{cnt: 0}

	h.SetLogCallback(func(ctx interface{}, lvl LogLevel, msg string) {
		cnt := ctx.(*Cnt)
		cnt.cnt++
	}, cnt)

	h.Release()

	if cnt.cnt != 1 {
		panic(nil)
	}
}

// +build six

package alpm

import (
	"fmt"
	"testing"
)

type ICnt interface {
	get() int
}

type Cnt struct {
	cnt int
}

func (c *Cnt) get() int {
	ret := c.cnt
	c.cnt += 1
	return ret
}

func TestCallbacks(t *testing.T) {
	h, _ := Initialize("/", "/var/lib/pacman")
	defer h.Release()

	h.SetLogCallback(func(ctx interface{}, lvl LogLevel, msg string) {
		cnt := ctx.(*Cnt)
		fmt.Printf("go-alpm: cnt(%d) %s\n", cnt.get(), msg)
	}, &Cnt{cnt: 0})

	h.SetQuestionCallback(func(ctx interface{}, question QuestionAny) {
		cnt := ctx.(*Cnt)
		fmt.Printf("question: cnt(%d)", cnt.get())
	}, &Cnt{cnt: 2})
}

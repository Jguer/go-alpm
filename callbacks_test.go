package alpm_test

import (
	"testing"

	alpm "github.com/Jguer/go-alpm/v2"
	"github.com/stretchr/testify/require"
)

type Cnt struct {
	cnt int
}

func TestCallbacks(t *testing.T) {
	t.Parallel()
	h, er := alpm.Initialize(root, dbpath)
	defer h.Release()
	if er != nil {
		t.Errorf("Failed at alpm initialization: %s", er)
	}

	cnt := &Cnt{cnt: 0}

	h.SetLogCallback(func(ctx interface{}, lvl alpm.LogLevel, msg string) {
		cnt := ctx.(*Cnt)
		cnt.cnt++
	}, cnt)

	h.Release()

	require.Equal(t, 1, cnt.cnt)
}

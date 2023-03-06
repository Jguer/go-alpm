package alpm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	alpm "github.com/Jguer/go-alpm/v2"
)

func TestPkgFilesList(t *testing.T) {
	t.Parallel()
	const (
		root   = "/"
		dbpath = "/var/lib/pacman"
	)

	h, er := alpm.Initialize(root, dbpath)
	defer h.Release()
	if er != nil {
		t.Errorf("Failed at alpm initialization: %s", er)
	}

	db, _ := h.LocalDB()

	pkg := db.Pkg("glibc")

	files := pkg.Files()
	assert.NotEmpty(t, files)
	assert.NotNil(t, files[0])
}

// alpm_test.go - Tests for alpm.go.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package alpm_test

import (
	"testing"

	alpm "github.com/Jguer/go-alpm/v2"
	"github.com/stretchr/testify/assert"
)

const (
	root   = "/"
	dbpath = "/var/lib/pacman"
)

func TestExampleVerCmp(t *testing.T) {
	t.Parallel()

	assert.Less(t, alpm.VerCmp("1.0-2", "2.0-1"), 0)
	assert.Equal(t, alpm.VerCmp("2.0.2-2", "2.0.2-2"), 0)
	assert.Greater(t, alpm.VerCmp("1:1.0-2", "2.0-1"), 0)
}

func TestRevdeps(t *testing.T) {
	t.Parallel()

	h, _ := alpm.Initialize(root, dbpath)

	db, _ := h.LocalDB()
	pkg := db.Pkg("glibc")
	for i, pkgname := range pkg.ComputeRequiredBy() {
		t.Logf(pkgname)
		if i == 10 {
			t.Logf("and %d more...", len(pkg.ComputeRequiredBy())-10)
			return
		}
	}
}

func TestLocalDB(t *testing.T) {
	t.Parallel()
	h, _ := alpm.Initialize(root, dbpath)

	defer func() {
		if recover() != nil {
			t.Errorf("local db failed")
		}
	}()
	db, _ := h.LocalDB()
	number := 0
	for _, pkg := range db.PkgCache().Slice() {
		number++
		if number <= 15 {
			t.Logf("%v", pkg.Name())
		}
	}
	if number > 15 {
		t.Logf("%d more packages...", number-15)
	}
}

func TestRelease(t *testing.T) {
	t.Parallel()
	h, _ := alpm.Initialize(root, dbpath)

	if err := h.Release(); err != nil {
		t.Error(err)
	}
}

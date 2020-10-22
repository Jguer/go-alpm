//
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package main

import (
	"fmt"

	"github.com/Jguer/go-alpm/v2"
)

func main() {
	h, er := alpm.Initialize("/", "/var/lib/pacman")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer h.Release()

	db, _ := h.RegisterSyncDB("core", 0)
	h.RegisterSyncDB("community", 0)
	h.RegisterSyncDB("extra", 0)

	for _, pkg := range db.PkgCache().Slice() {
		fmt.Printf("%s %s\n  %s\n",
			pkg.Name(), pkg.Version(), pkg.Description())
	}
}

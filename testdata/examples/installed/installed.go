// installed.go - Example of getting a list of installed packages.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package main

import (
	"fmt"
	"os"

	"github.com/Jguer/go-alpm/v2"
)

func main() {
	h, er := alpm.Initialize("/", "/var/lib/pacman")
	if er != nil {
		print(er, "\n")
		os.Exit(1)
	}

	db, er := h.LocalDB()
	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	}

	for _, pkg := range db.PkgCache().Slice() {
		fmt.Printf("%s %s\n", pkg.Name(), pkg.Version())
	}

	if h.Release() != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

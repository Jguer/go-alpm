// sync.go - Example of installing, and removing packages.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package main

import (
	"fmt"
	"log"

	"github.com/Jguer/go-alpm/v2"
	"github.com/Morganamilo/go-pacmanconf"
)

func main() {
	h, err := alpm.Initialize("/", "/var/lib/pacman")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := h.Release(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	conf, _, err := pacmanconf.ParseFile("/etc/pacman.conf")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("loading databases")

	var addPkg alpm.IPackage
	var removePkg alpm.IPackage

	for _, repo := range conf.Repos {
		db, err := h.RegisterSyncDB(repo.Name, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		db.SetServers(repo.Servers)

		foundPkg := db.Pkg("vim")
		if foundPkg != nil {
			addPkg = foundPkg
		}
	}

	localDb, err := h.LocalDB()
	if err != nil {
		log.Println(err)
		return
	}
	// A package can only be removed if it is in the local DB
	removePkg = localDb.Pkg("vi")

	fmt.Println("initializing transaction")
	if err := h.TransInit(0); err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		fmt.Println("releasing transaction")
		if err := h.TransRelease(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	fmt.Println("synchronizing databases")
	if err := h.SyncSysupgrade(true); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("adding package")
	if err := h.AddPkg(addPkg); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("removing package")
	if removePkg != nil {
		if err := h.RemovePkg(removePkg); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("vi is not installed")
	}

	fmt.Println("preparing transaction")
	if err := h.TransPrepare(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("committing transaction")
	if err := h.TransCommit(); err != nil {
		fmt.Println(err)
		return
	}
}

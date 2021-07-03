//
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package main

import (
	"fmt"
	"log"

	paconf "github.com/Morganamilo/go-pacmanconf"
)

func human(size int64) string {
	floatsize := float32(size)
	units := [...]string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi", "Yi"}
	for _, unit := range units {
		if floatsize < 1024 {
			return fmt.Sprintf("%.1f %sB", floatsize, unit)
		}
		floatsize /= 1024
	}
	return fmt.Sprintf("%d%s", size, "B")
}

func upgrades(h *alpm.Handle) ([]alpm.IPackage, error) {
	localDb, err := h.LocalDB()
	if err != nil {
		return nil, err
	}

	syncDbs, err := h.SyncDBs()
	if err != nil {
		return nil, err
	}

	slice := []alpm.IPackage{}
	for _, pkg := range localDb.PkgCache().Slice() {
		newPkg := pkg.SyncNewVersion(syncDbs)
		if newPkg != nil {
			slice = append(slice, newPkg)
		}
	}
	return slice, nil
}

func main() {
	h, er := alpm.Initialize("/", "/var/lib/pacman")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer h.Release()

	PacmanConfig, _, err := paconf.ParseFile("/etc/pacman.conf")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
		We have to configure alpm with pacman configuration
		to load the repositories and other stuff
	*/
	for _, repo := range PacmanConfig.Repos {
		db, err := h.RegisterSyncDB(repo.Name, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		db.SetServers(repo.Servers)

		/*
			Configure repository usage to match with
			the alpm library provided formats
		*/
		if len(repo.Usage) == 0 {
			db.SetUsage(alpm.UsageAll)
		}
		for _, usage := range repo.Usage {
			switch usage {
			case "Sync":
				db.SetUsage(alpm.UsageSync)
			case "Search":
				db.SetUsage(alpm.UsageSearch)
			case "Install":
				db.SetUsage(alpm.UsageInstall)
			case "Upgrade":
				db.SetUsage(alpm.UsageUpgrade)
			case "All":
				db.SetUsage(alpm.UsageAll)
			}
		}
	}

	upgrades, err := upgrades(h)
	if err != nil {
		log.Fatalln(err)
	}

	var size int64 = 0
	for _, pkg := range upgrades {
		size += pkg.Size()
		fmt.Printf("%s %s -> %s\n", pkg.Name(), pkg.Version(),
			pkg.Version())
	}
	fmt.Printf("Total Download Size: %s\n", human(size))
}

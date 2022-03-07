//go:build !next
// +build !next

// db.go - Functions for database handling.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package alpm

/*
#include <alpm.h>
#include <alpm_list.h>
*/
import "C"

import (
	"unsafe"
)

// PkgCachebyGroup returns a PackageList of packages belonging to a group
func (l DBList) FindGroupPkgs(name string) IPackageList {
	cName := C.CString(name)

	defer C.free(unsafe.Pointer(cName))

	pkglist := (*C.struct___alpm_list_t)(unsafe.Pointer(l.list))
	pkgcache := (*list)(unsafe.Pointer(C.alpm_find_group_pkgs(pkglist, cName)))

	return PackageList{pkgcache, l.handle}
}

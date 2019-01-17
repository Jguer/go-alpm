// db.go - Functions for database handling.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package alpm

/*
#include <alpm.h>
*/
import "C"

import "unsafe"

// PkgCachebyGroup returns a PackageList of packages belonging to a group
func (l DBList) FindGroupPkgs(name string) PackageList {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	pkglist := unsafe.Pointer(l.List)

	pkgcache := C.alpm_find_group_pkgs((*C.alpm_list_t)(pkglist), cName)
	if pkgcache == nil {
		return makePackageList(pkgcache, l.handle)
	}

	return makePackageList(pkgcache, l.handle)
}

// NewVersion checks if there is a new version of the package in a given DBlist.
func (pkg *Package) SyncNewVersion(l DBList) *Package {
	ptr := C.alpm_sync_get_new_version(pkg.pmpkg,
		(*C.alpm_list_t)(unsafe.Pointer(l.List)))
	if ptr == nil {
		return nil
	}
	return &Package{ptr, l.handle}
}

func (h *Handle) SyncSysupgrade(enableDowngrade bool) error {
	intEnableDowngrade := C.int(0)

	if enableDowngrade {
		intEnableDowngrade = C.int(1)
	}

	ret := C.alpm_sync_sysupgrade(h.ptr, intEnableDowngrade)
	if ret != 0 {
		return h.LastError()
	}

	return nil
}

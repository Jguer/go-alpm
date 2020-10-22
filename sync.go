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

// SyncNewVersion checks if there is a new version of the
// package in a given DBlist.
func (pkg *Package) SyncNewVersion(l IDBList) IPackage {
	ptr := C.alpm_sync_get_new_version(pkg.pmpkg,
		(*C.alpm_list_t)(unsafe.Pointer(l.(*DBList).list)))
	if ptr == nil {
		return nil
	}
	return &Package{ptr, l.(*DBList).handle}
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

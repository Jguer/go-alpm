//go:build !next
// +build !next

// package.go - libalpm package type and methods.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package alpm

/*
#include <alpm.h>

int pkg_cmp(const void *v1, const void *v2)
{
    alpm_pkg_t *p1 = (alpm_pkg_t *)v1;
    alpm_pkg_t *p2 = (alpm_pkg_t *)v2;
    off_t s1 = alpm_pkg_get_isize(p1);
    off_t s2 = alpm_pkg_get_isize(p2);

    if (s1 > s2)
        return -1;
    else if (s1 < s2)
        return 1;
    else
        return 0;
}
*/
import "C"

import (
	"unsafe"
)

// SortBySize returns a PackageList sorted by size.
func (l PackageList) SortBySize() IPackageList {
	pkgList := (*C.struct___alpm_list_t)(unsafe.Pointer(l.list))

	pkgCache := (*list)(unsafe.Pointer(
		C.alpm_list_msort(pkgList, //nolint
			C.alpm_list_count(pkgList),
			C.alpm_list_fn_cmp(C.pkg_cmp))))
	if pkgCache == nil {
		return nil
	}

	return PackageList{pkgCache, l.handle}
}

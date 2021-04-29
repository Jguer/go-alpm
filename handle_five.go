// +build !six
// handle.go - libalpm handle type and methods.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

// Package alpm implements Go bindings to the libalpm library used by Pacman,
// the Arch Linux package manager. Libalpm allows the creation of custom front
// ends to the Arch Linux package ecosystem.
//
// Libalpm does not include support for the Arch User Repository (AUR).
package alpm

// #include <alpm.h>
import "C"

func (h *Handle) Arch() (string, error) {
	return h.optionGetStr(func(handle *C.alpm_handle_t) *C.char {
		return C.alpm_option_get_arch(handle)
	})
}

func (h *Handle) SetArch(str string) error {
	return h.optionSetStr(str, func(handle *C.alpm_handle_t, cStr *C.char) C.int {
		return C.alpm_option_set_arch(handle, cStr)
	})
}

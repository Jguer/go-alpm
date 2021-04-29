// +build six
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

func (h *Handle) GetArchitectures() (StringList, error) {
	return h.optionGetList(func(handle *C.alpm_handle_t) *C.alpm_list_t {
		return C.alpm_option_get_architectures(handle)
	})
}

func (h *Handle) SetArchitectures(str []string) error {
	return h.optionSetList(str, func(handle *C.alpm_handle_t, l *C.alpm_list_t) C.int {
		return C.alpm_option_set_architectures(handle, l)
	})
}

func (h *Handle) AddArchitecture(str string) error {
	return h.optionAddList(str, func(handle *C.alpm_handle_t, cStr *C.char) C.int {
		return C.alpm_option_add_architecture(handle, cStr)
	})
}

func (h *Handle) RemoveArchitecture(str string) (bool, error) {
	return h.optionRemoveList(str, func(handle *C.alpm_handle_t, cStr *C.char) C.int {
		return C.alpm_option_remove_architecture(handle, cStr)
	})
}

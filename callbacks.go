// callbacks.go - Handles libalpm callbacks.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package alpm

/*
#include "callbacks.h"
*/
import "C"
import "unsafe"

var DefaultLogLevel = LogWarning

type (
	logCallbackSig      func(interface{}, LogLevel, string)
	questionCallbackSig func(interface{}, QuestionAny)
	callbackContextPool map[C.go_ctx_t]interface{}
)

var (
	logCallbackContextPool      callbackContextPool = callbackContextPool{}
	questionCallbackContextPool callbackContextPool = callbackContextPool{}
)

func DefaultLogCallback(ctx interface{}, lvl LogLevel, s string) {
	if lvl <= DefaultLogLevel {
		print("go-alpm: ", s)
	}
}

//export go_alpm_go_log_callback
func go_alpm_go_log_callback(goCb unsafe.Pointer, goCtx C.go_ctx_t, lvl C.alpm_loglevel_t, s *C.char) {
	cb := *(*logCallbackSig)(goCb)
	ctx := logCallbackContextPool[goCtx]

	cb(ctx, LogLevel(lvl), C.GoString(s))
}

//export go_alpm_go_question_callback
func go_alpm_go_question_callback(goCb unsafe.Pointer, goCtx C.go_ctx_t, question *C.alpm_question_t) {
	q := (*C.alpm_question_any_t)(unsafe.Pointer(question))

	cb := *(*questionCallbackSig)(goCb)
	ctx := questionCallbackContextPool[goCtx]

	cb(ctx, QuestionAny{q})
}

func (h *Handle) SetLogCallback(cb logCallbackSig, ctx interface{}) {
	goCb := unsafe.Pointer(&cb)
	goCtx := C.go_ctx_t(h.ptr)

	logCallbackContextPool[goCtx] = ctx

	C.go_alpm_set_log_callback(h.ptr, goCb, goCtx)
}

func (h *Handle) SetQuestionCallback(cb questionCallbackSig, ctx interface{}) {
	goCb := unsafe.Pointer(&cb)
	goCtx := C.go_ctx_t(h.ptr)

	questionCallbackContextPool[goCtx] = ctx

	C.go_alpm_set_question_callback(h.ptr, goCb, goCtx)
}

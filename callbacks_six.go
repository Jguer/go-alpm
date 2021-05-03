// +build six

package alpm

/*
#cgo CFLAGS: -DSIX=1
#include "callbacks_six.h"
*/
import "C"

import (
	"unsafe"
)

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
func go_alpm_go_log_callback(go_cb unsafe.Pointer, go_ctx C.go_ctx_t, lvl C.alpm_loglevel_t, s *C.char) {
	cb := *(*logCallbackSig)(go_cb)
	ctx := logCallbackContextPool[go_ctx]

	cb(ctx, LogLevel(lvl), C.GoString(s))
}

//export go_alpm_go_question_callback
func go_alpm_go_question_callback(go_cb unsafe.Pointer, go_ctx C.go_ctx_t, question *C.alpm_question_t) {
	q := (*C.alpm_question_any_t)(unsafe.Pointer(question))

	cb := *(*questionCallbackSig)(go_cb)
	ctx := questionCallbackContextPool[go_ctx]

	cb(ctx, QuestionAny{q})
}

func (h *Handle) SetLogCallback(cb logCallbackSig, ctx interface{}) {
	go_cb := unsafe.Pointer(&cb)
	go_ctx := C.go_ctx_t(h.ptr)

	logCallbackContextPool[go_ctx] = ctx

	C.go_alpm_set_log_callback(h.ptr, go_cb, go_ctx)
}

func (h *Handle) SetQuestionCallback(cb questionCallbackSig, ctx interface{}) {
	go_cb := unsafe.Pointer(&cb)
	go_ctx := C.go_ctx_t(h.ptr)

	questionCallbackContextPool[go_ctx] = ctx

	C.go_alpm_set_question_callback(h.ptr, go_cb, go_ctx)
}

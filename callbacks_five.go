// +build !six

package alpm

/*
#cgo CFLAGS: -DSIX=0
#include <alpm.h>
void go_alpm_set_question(alpm_handle_t *handle);
void go_alpm_set_logging(alpm_handle_t *handle);
*/
import "C"

import (
	"unsafe"
)

type (
	logCallbackSig      func(LogLevel, string)
	questionCallbackSig func(QuestionAny)
)

func DefaultLogCallback(lvl LogLevel, s string) {
	if lvl <= DefaultLogLevel {
		print("go-alpm: ", s)
	}
}

var (
	globalLogCallback      logCallbackSig
	globalQuestionCallback questionCallbackSig
)

//export logCallback
func logCallback(level C.alpm_loglevel_t, cstring *C.char) {
	globalLogCallback(LogLevel(level), C.GoString(cstring))
}

//export questionCallback
func questionCallback(question *C.alpm_question_t) {
	q := (*C.alpm_question_any_t)(unsafe.Pointer(question))
	globalQuestionCallback(QuestionAny{q})
}

func (h *Handle) SetLogCallback(cb logCallbackSig) {
	globalLogCallback = cb
	C.go_alpm_set_logging(h.ptr)
}

func (h *Handle) alpmSetQuestion(cb questionCallbackSig) {
	globalQuestionCallback = cb
	C.go_alpm_set_question(h.ptr)
}

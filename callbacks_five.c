// callbacks.c - Sets alpm callbacks to Go functions.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

#if !SIX

#include <stdint.h>
#include <stdio.h>
#include <stdarg.h>
#include <alpm.h>

void logCallback(alpm_loglevel_t level, char *cstring);
void questionCallback(alpm_question_t *question);

static void _log_cb(alpm_loglevel_t level, const char *fmt, va_list arg) {
  char *s = malloc(128);
  if (s == NULL) return;
  int16_t length = vsnprintf(s, 128, fmt, arg);
  if (length > 128) {
    length = (length + 16) & ~0xf;
    s = realloc(s, length);
  }
  if (s != NULL) {
		logCallback(level, s);
		free(s);
  }
}

static void _question_cb(alpm_question_t *question) {
	questionCallback(question);
}

void go_alpm_set_logging(alpm_handle_t *handle) {
	alpm_option_set_logcb(handle, _log_cb);
}

void go_alpm_set_question(alpm_handle_t *handle) {
	alpm_option_set_questioncb(handle, _question_cb);
}

#endif

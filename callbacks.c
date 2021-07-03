// callbacks.c - Sets alpm callbacks to Go functions.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

#include <stdint.h>
#include <stdio.h>
#include <stdarg.h>
#include "callbacks.h"

typedef struct {
  void *go_cb;
  go_ctx_t go_ctx;
} go_alpm_context_t;

static void _log_cb(void *c_ctx, alpm_loglevel_t level, const char *fmt, va_list arg) {
  go_alpm_context_t *ctx = c_ctx;
  char *s = malloc(128);
  if (s == NULL) return;
  int16_t length = vsnprintf(s, 128, fmt, arg);
  if (length > 128) {
    length = (length + 16) & ~0xf;
    s = realloc(s, length);
  }
  if (s != NULL) {
    go_alpm_go_log_callback(&ctx->go_cb, ctx->go_ctx, level, s);
		free(s);
  }
}

static void _question_cb(void *c_ctx, alpm_question_t *question) {
  go_alpm_context_t *ctx = c_ctx;
  go_alpm_go_question_callback(&ctx->go_cb, ctx->go_ctx, question);
}

static void *alloc_ctx(go_alpm_context_t *ctx, void *go_cb, go_ctx_t go_ctx) {
  if (ctx == NULL ) {
    ctx = malloc(sizeof(go_alpm_context_t));
  }

  ctx->go_cb = go_cb;
  ctx->go_ctx = go_ctx;

  return ctx;
}

void go_alpm_set_log_callback(alpm_handle_t *handle, void *go_cb, go_ctx_t go_ctx) {
  void *ctx = alpm_option_get_logcb_ctx(handle);
  ctx = alloc_ctx(ctx, *(void**)go_cb, go_ctx);
  alpm_option_set_logcb(handle, _log_cb, ctx);
}

void go_alpm_set_question_callback(alpm_handle_t *handle, void *go_cb, go_ctx_t go_ctx) {
  void *ctx = alpm_option_get_questioncb_ctx(handle);
  ctx = alloc_ctx(ctx, *(void**)go_cb, go_ctx);
  alpm_option_set_questioncb(handle, _question_cb, ctx);
}

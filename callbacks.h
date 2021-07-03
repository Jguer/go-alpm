#include <alpm.h>

typedef void *go_ctx_t;

void go_alpm_go_log_callback(void *go_cb, go_ctx_t go_ctx, alpm_loglevel_t level, char *message);
void go_alpm_go_question_callback(void *go_cb, go_ctx_t go_ctx, alpm_question_t *question);

void go_alpm_set_log_callback(alpm_handle_t *handle, void *go_cb, go_ctx_t go_ctx);
void go_alpm_set_question_callback(alpm_handle_t *handle, void *go_cb, go_ctx_t go_ctx);

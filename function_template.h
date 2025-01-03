#ifndef V8GO_FUNCTION_TEMPLATE_H
#define V8GO_FUNCTION_TEMPLATE_H

#include "errors.h"

#ifdef __cplusplus

namespace v8 {
class Isolate;
}

typedef v8::Isolate v8Isolate;

extern "C" {
#else

typedef struct v8Isolate v8Isolate;

#endif

typedef struct m_template m_template;
typedef struct m_ctx m_ctx;

extern m_template* NewFunctionTemplate(v8Isolate* iso_ptr, int callback_ref);
extern RtnValue FunctionTemplateGetFunction(m_template* ptr, m_ctx* ctx_ptr);

#ifdef __cplusplus
}  // extern "C"
#endif
#endif
#ifndef V8GO_SYMBOL_H
#define V8GO_SYMBOL_H

#include "forward-declarations.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct m_value m_value;
typedef m_value* ValuePtr;

typedef enum {
  SYMBOL_ASYNC_ITERATOR = 1,
  SYMBOL_HAS_INSTANCE,
  SYMBOL_IS_CONCAT_SPREADABLE,
  SYMBOL_ITERATOR,
  SYMBOL_MATCH,
  SYMBOL_REPLACE,
  SYMBOL_SEARCH,
  SYMBOL_SPLIT,
  SYMBOL_TO_PRIMITIVE,
  SYMBOL_TO_STRING_TAG,
  SYMBOL_UNSCOPABLES,
} SymbolIndex;

ValuePtr BuiltinSymbol(IsolatePtr iso_ptr, SymbolIndex idx);
const char* SymbolDescription(ValuePtr ptr);

#ifdef __cplusplus
}
#endif
#endif

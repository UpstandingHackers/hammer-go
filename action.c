#include <assert.h>
#include "hammer/src/parsers/parser_internal.h"
#include "action.h"

// prototypes
HParser* go_action__m(HAllocator* mm__, const HParser* p, void* a);

typedef struct {
  const HParser *p;
  // action is a Go *ActionFunc
  void* action;
} GoParseAction;

static HParseResult* parse_action(void *env, HParseState *state) {
  GoParseAction *a = (GoParseAction*)env;
  if (a->p && a->action) {
    HParseResult *tmp = h_do_parse(a->p, state);
    if(tmp) {
        const HParsedToken *tok = go_action_hook(a->action, tmp);
        return make_result(state->arena, (HParsedToken*)tok);
    } else
      return NULL;
  } else // either the parser's missing or the action's missing
    return NULL;
}

static const HParserVtable action_vt = {
  .parse = parse_action,
};

HParser* go_action(const HParser* p, void* a) {
  return go_action__m(&system_allocator, p, a);
}

HParser* go_action__m(HAllocator* mm__, const HParser* p, void* a) {
  GoParseAction *env = h_new(GoParseAction, 1);
  env->p = p;
  env->action = a;

  return h_new_parser(mm__, &action_vt, env);
}

void assignUintValue(HParsedToken* token, uint64_t value) {
  token->uint = value;
}

void assignVoidValue(HParsedToken* token, void* value) {
  token->user = value;
}

#include <assert.h>
#include <hammer/parsers/parser_internal.h>
#include "action.h"

// prototypes
HParser* go_action__m(HAllocator* mm__, const HParser* p, void* a);

typedef struct {
	const HParser *p;
 	void* action; // Go type *ActionFunc
} GoParseAction;

static HParseResult* parse_action(void *env, HParseState *state) {
	GoParseAction *a = (GoParseAction*)env;
	if (a->p == NULL || a->action == NULL) {
		// Bad parser. This shouldn't ever happen.
		return NULL;
	}

	HParseResult *pr = h_do_parse(a->p, state);
	if (pr == NULL) {
		// Upstream HParser failed
		return NULL;
	}

	GoActionResult gar = go_action_hook(a->action, pr);
	if (!gar.valid) {
		// ActionFunc declared failed parse
		return NULL;
	}

	return make_result(state->arena, gar.token);
}

static void desugar_action(HAllocator *mm__, HCFStack *stk__, void *env) {
	GoParseAction *a = (GoParseAction*)env;

	HCFS_BEGIN_CHOICE() {
		HCFS_BEGIN_SEQ() {
			HCFS_DESUGAR(a->p);
		} HCFS_END_SEQ();
		HCFS_THIS_CHOICE->action = a->action;
		HCFS_THIS_CHOICE->reshape = h_act_first;
	} HCFS_END_CHOICE();
}

static bool action_isValidRegular(void *env) {
	return false;
}

static bool action_isValidCF(void *env) {
	GoParseAction *a = (GoParseAction*)env;
	return a->p->vtable->isValidCF(a->p->env);
}

static const HParserVtable action_vt = {
	.parse = parse_action,
	.isValidRegular = action_isValidRegular,
	.isValidCF = action_isValidCF,
	.desugar = desugar_action,
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

HTokenType getGoTokenId() {
	return h_allocate_token_type("com.prevoty.gotoken");
}


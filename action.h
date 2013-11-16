#pragma once

#include <hammer.h>

// implemented in action.go
extern HParsedToken* go_action_hook(void*, HParseResult*);

// implemented in action.c
HParser* go_action(const HParser* p, void* a);

void assignUintValue(HParsedToken* token, uint64_t value);
void assignVoidValue(HParsedToken* token, void* value);
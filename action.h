#pragma once

#include <hammer/hammer.h>


typedef struct {
  HParsedToken* token;
  bool valid;
} GoActionResult;


// implemented in action.go
extern GoActionResult go_action_hook(void*, HParseResult*);

// implemented in action.c
HParser* go_action(const HParser* p, void* a);
void assignUintValue(HParsedToken* token, uint64_t value);
void assignVoidValue(HParsedToken* token, void* value);

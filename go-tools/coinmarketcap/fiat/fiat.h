#ifndef FITA_H
#define FITA_H

#include <tools.h>

DEFINE_OPTIONAL(String)

DEFINE_OPTIONAL(Int)

DEFINE_OPTIONAL(Bool)

typedef struct {
    Int id;
    String name;
    String sign;
    String symbol;
} Fiat;

void release_Fiat(Fiat data);

DEFINE_LIST(Fiat);

DEFINE_OPTIONAL(List_Fiat);

DEFINE_RESULT(List_Fiat);

#endif
#ifndef CONVERSION_TOOLS_H
#define CONVERSION_TOOLS_H

#include <tools.h>

DEFINE_OPTIONAL(String)

typedef struct {
    Float price;
    String last_updated;
} Quote;

void release_Quote(Quote data);

DEFINE_DICT(Quote)

typedef struct {
    Int id;
    String symbol;
    String name;
    Float amount;
    String last_updated;
    Dict_Quote quote;
} PriceConversion;

void release_PriceConversion(PriceConversion data);

DEFINE_OPTIONAL(PriceConversion)

DEFINE_RESULT(PriceConversion)

#endif

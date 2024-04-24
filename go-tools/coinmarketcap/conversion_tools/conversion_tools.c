#include <tools.h>
#include <conversion_tools.h>

IMPL_OPTIONAL(String)

void release_Quote(Quote data) {
    free(data.last_updated);
}

IMPL_DICT(Quote)

void release_PriceConversion(PriceConversion data) {
    free(data.symbol);
    free(data.name);
    free(data.last_updated);
    release_Dict_Quote(data.quote);
}

IMPL_OPTIONAL(PriceConversion)

IMPL_RESULT(PriceConversion)
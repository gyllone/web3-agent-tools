#include <tools.h>
#include <fiat.h>

IMPL_OPTIONAL(String)

IMPL_OPTIONAL(Int)

IMPL_OPTIONAL(Bool)

void release_Fiat(Fiat data) {
    free(data.name);
    free(data.sign);
    free(data.symbol);
}

IMPL_LIST(Fiat)

IMPL_OPTIONAL(List_Fiat)

IMPL_RESULT(List_Fiat)
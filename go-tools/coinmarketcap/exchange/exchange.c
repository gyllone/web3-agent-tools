#include <tools.h>
#include <exchange.h>

IMPL_OPTIONAL(String)

IMPL_OPTIONAL(Int)

IMPL_OPTIONAL(Bool)

void release_Exchange(Exchange data) {
    free(data.name);
    free(data.slug);
    free(data.first_historical_data);
    free(data.last_historical_data);
}

IMPL_LIST(Exchange)

IMPL_RESULT(List_Exchange)

IMPL_LIST(String)

void release_Metadata(Metadata data) {
    free(data.name);
    free(data.slug);
    free(data.description);
    free(data.notice);
    release_List_String(data.fiats);
    free(data.urls);
    free(data.date_launched);
    free(data.spot_volume_last_updated);
}

IMPL_DICT(Metadata)

IMPL_RESULT(Dict_Metadata)



void release_Platform(Platform data) {
    free(data.symbol);
    free(data.name);
}

void release_Currency(Currency data) {
    free(data.symbol);
    free(data.name);
}

void release_Asset(Asset data) {
    free(data.wallet_address);
    release_Platform(data.platform);
    release_Currency(data.currency);
}

IMPL_LIST(Asset)

IMPL_RESULT(List_Asset)

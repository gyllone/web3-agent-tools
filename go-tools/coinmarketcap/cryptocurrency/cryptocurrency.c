#include <tools.h>
#include <cryptocurrency.h>

IMPL_LIST(Float)

IMPL_LIST(List_Float)

IMPL_RESULT(List_List_Float)

IMPL_DICT(String)

IMPL_LIST(Dict_String)

IMPL_RESULT(List_Dict_String)

IMPL_DICT(Dict_String)

void release_MarketData(MarketData data) {
    release_Dict_String(data.metadata);
    release_Dict_Dict_String(data.quotes);
}

IMPL_LIST(MarketData)

IMPL_RESULT(List_MarketData)

void release_Category(Category data) {
    free(data.id);
    free(data.name);
    free(data.description);
    free(data.last_updated);
}

IMPL_LIST(Category)

IMPL_RESULT(List_Category)

void release_Quote(Quote data) {
    free(data.last_updated);
}

IMPL_DICT(Quote)

void release_Coin(Coin data) {
    free(data.name);
    free(data.symbol);
    free(data.slug);
    free(data.date_added);
    free(data.tags);
    free(data.last_updated);
    release_Dict_Quote(data.quote);
}

IMPL_LIST(Coin)

void release_CategorySingle(CategorySingle data) {
    free(data.id);
    free(data.name);
    free(data.description);
    free(data.last_updated);
    release_List_Coin(data.coins);
}

IMPL_OPTIONAL(CategorySingle)

IMPL_RESULT(Optional_CategorySingle)

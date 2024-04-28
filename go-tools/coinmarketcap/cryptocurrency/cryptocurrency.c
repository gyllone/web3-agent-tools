#include <tools.h>
#include <cryptocurrency.h>

//IMPL_LIST(Float)
//
//IMPL_LIST(List_Float)
//
//IMPL_RESULT(List_List_Float)

IMPL_OPTIONAL(String)

IMPL_OPTIONAL(Int)

IMPL_OPTIONAL(Bool)

void release_Quote(Quote data) {
    free(data.last_updated);
}

IMPL_DICT(Quote)

void release_Platform(Platform data) {
    free(data.name);
    free(data.symbol);
    free(data.slug);
    free(data.token_address);
}

IMPL_OPTIONAL(Platform)

void release_Tag(Tag data) {
    free(data.slug);
    free(data.name);
    free(data.category);
}

IMPL_LIST(Tag)

void release_QuoteData(QuoteData data) {
    free(data.name);
    free(data.symbol);
    free(data.slug);
    free(data.date_added);
    release_List_Tag(data.tags);
    release_Optional_Platform(data.platform);
    free(data.last_updated);
    release_Dict_Quote(data.quote);
}

IMPL_DICT(QuoteData)

IMPL_OPTIONAL(Dict_QuoteData)

IMPL_RESULT(Dict_QuoteData)



void release_Cryptocurrency(Cryptocurrency data) {
    free(data.name);
    free(data.symbol);
    free(data.slug);
    free(data.first_historical_data);
    free(data.last_historical_data);
    release_Optional_Platform(data.platform);
}

IMPL_LIST(Cryptocurrency)

IMPL_OPTIONAL(List_Cryptocurrency)

IMPL_RESULT(List_Cryptocurrency)





IMPL_LIST(String)

void release_URLs(URLs data) {
    release_List_String(data.website);
    release_List_String(data.twitter);
    release_List_String(data.message_board);
    release_List_String(data.chat);
    release_List_String(data.facebook);
    release_List_String(data.explorer);
    release_List_String(data.reddit);
    release_List_String(data.technical_doc);
    release_List_String(data.source_code);
    release_List_String(data.announcement);
}

void release_Metadata(Metadata data) {
    free(data.name);
    free(data.symbol);
    free(data.category);
    free(data.description);
    free(data.slug);
    free(data.logo);
    free(data.subreddit);
    free(data.notice);
    release_List_String(data.tags);
    release_List_String(data.tag_names);
    release_List_String(data.tag_groups);
    release_URLs(data.urls);
    release_Optional_Platform(data.platform);
    free(data.date_added);
    free(data.twitter_username);
    free(data.date_launched);
}

IMPL_DICT(Metadata)

IMPL_OPTIONAL(Dict_Metadata)

IMPL_RESULT(Dict_Metadata)




void release_MarketData(MarketData data) {
    free(data.name);
    free(data.symbol);
    free(data.slug);
    free(data.date_added);
    release_List_String(data.tags);
    release_Optional_Platform(data.platform);
    free(data.last_updated);
    release_Dict_Quote(data.quote);
}

IMPL_LIST(MarketData)

IMPL_OPTIONAL(List_MarketData)

IMPL_RESULT(List_MarketData)




void release_Category(Category data) {
    free(data.id);
    free(data.name);
    free(data.description);
    free(data.last_updated);
}

IMPL_LIST(Category)

IMPL_OPTIONAL(List_Category)

IMPL_RESULT(List_Category)




void release_Coin(Coin data) {
    free(data.name);
    free(data.symbol);
    free(data.slug);
    free(data.date_added);
    release_List_String(data.tags);
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

IMPL_RESULT(CategorySingle)

#ifndef CRYPTOCURRENCY_H
#define CRYPTOCURRENCY_H

#include <tools.h>

DEFINE_OPTIONAL(String)

DEFINE_OPTIONAL(Int)

DEFINE_OPTIONAL(Bool)


typedef struct {
    String last_updated;
    Float price;
    Float volume_24h;
    Float volume_change_24h;
    Float percent_change_1h;
    Float percent_change_24h;
    Float percent_change_7d;
    Float percent_change_30d;
    Float percent_change_60d;
    Float percent_change_90d;
    Float market_cap;
    Float market_cap_dominance;
    Float fully_diluted_market_cap;
    Float tvl;
} Quote;

void release_Quote(Quote data);

DEFINE_DICT(Quote)

typedef struct {
    Int id;
    String name;
    String symbol;
    String slug;
    String token_address;
} Platform;

void release_Platform(Platform data);

DEFINE_OPTIONAL(Platform)

typedef struct {
    String slug;
    String name;
    String category;
} Tag;

void release_Tag(Tag data);

DEFINE_LIST(Tag)

typedef struct {
    Int id;
    String name;
    String symbol;
    String slug;
    Int num_market_pairs;
    String date_added;
    List_Tag tags;
    Int max_supply;
    Float circulating_supply;
    Float total_supply;
    Int is_active;
    Bool infinite_supply;
    Optional_Platform platform;
    Int cmc_rank;
    Int is_fiat;
    Float self_reported_circulating_supply;
    Float self_reported_market_cap;
    Float tvl_ratio;
    String last_updated;
    Dict_Quote quote;
} QuoteData;

void release_QuoteData(QuoteData data);

DEFINE_DICT(QuoteData)

DEFINE_OPTIONAL(Dict_QuoteData)

DEFINE_RESULT(Dict_QuoteData)



typedef struct {
    Int id;
    Int rank;
    String name;
    String symbol;
    String slug;
    Int is_active;
    String first_historical_data;
    String last_historical_data;
    Optional_Platform platform;
} Cryptocurrency;

void release_Cryptocurrency(Cryptocurrency data);

DEFINE_LIST(Cryptocurrency)

DEFINE_OPTIONAL(List_Cryptocurrency)

DEFINE_RESULT(List_Cryptocurrency)




DEFINE_LIST(String)

typedef struct {
    List_String website;
    List_String twitter;
    List_String message_board;
    List_String chat;
    List_String facebook;
    List_String explorer;
    List_String reddit;
    List_String technical_doc;
    List_String source_code;
    List_String announcement;
} URLs;

void release_URLs(URLs data);

typedef struct {
    Int id;
    String name;
    String symbol;
    String category;
    String description;
    String slug;
    String logo;
    String subreddit;
    String notice;
    List_String tags;
    List_String tag_names;
    List_String tag_groups;
    URLs urls;
    Optional_Platform platform;
    String date_added;
    String twitter_username;
    Int is_hidden;
    String date_launched;
    Float self_reported_circulating_supply;
    Float self_reported_market_cap;
    Bool infinite_supply;
} Metadata;

void release_Metadata(Metadata data);

DEFINE_DICT(Metadata)

DEFINE_OPTIONAL(Dict_Metadata)

DEFINE_RESULT(Dict_Metadata)



typedef struct {
    Int id;
    String name;
    String symbol;
    String slug;
    Int num_market_pairs;
    String date_added;
    List_String tags;
    Int max_supply;
    Int circulating_supply;
    Int total_supply;
    Bool infinite_supply;
    Optional_Platform platform;
    Int cmc_rank;
    Float self_reported_circulating_supply;
    Float self_reported_market_cap;
    Float tvl_ratio;
    String last_updated;
    Dict_Quote quote;
} MarketData;

void release_MarketData(MarketData data);

DEFINE_LIST(MarketData)

DEFINE_OPTIONAL(List_MarketData)

DEFINE_RESULT(List_MarketData)



typedef struct {
    String id;
    String name;
    String description;
    Int num_tokens;
    Float avg_price_change;
    Float market_cap;
    Float market_cap_change;
    Float volume;
    Float volume_change;
    String last_updated;
} Category;

void release_Category(Category data);

DEFINE_LIST(Category)

DEFINE_OPTIONAL(List_Category)

DEFINE_RESULT(List_Category)



typedef struct {
    Int id;
    String name;
    String symbol;
    String slug;
    Int num_market_pairs;
    String date_added;
    List_String tags;
    Int max_supply;
    Int circulating_supply;
    Int total_supply;
    Int is_active;
    Bool infinite_supply;
    Int cmc_rank;
    Int is_fiat;
    Float tvl_ratio;
    String last_updated;
    Dict_Quote quote;
} Coin;

void release_Coin(Coin data);

DEFINE_LIST(Coin)

typedef struct {
    String id;
    String name;
    String description;
    Int num_tokens;
    Float avg_price_change;
    Float market_cap;
    Float market_cap_change;
    Float volume;
    Float volume_change;
    String last_updated;
    List_Coin coins;
} CategorySingle;

void release_CategorySingle(CategorySingle data);

DEFINE_OPTIONAL(CategorySingle)

DEFINE_RESULT(CategorySingle)
#endif
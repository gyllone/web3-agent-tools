#ifndef CRYPTOCURRENCY_H
#define CRYPTOCURRENCY_H

#include <tools.h>

DEFINE_OPTIONAL(String)

DEFINE_OPTIONAL(Int)

DEFINE_OPTIONAL(Bool)

DEFINE_LIST(Float)

DEFINE_LIST(List_Float)

DEFINE_RESULT(List_List_Float)

DEFINE_DICT(String)

DEFINE_LIST(Dict_String)

DEFINE_RESULT(List_Dict_String)

DEFINE_DICT(Dict_String)

typedef struct {
    Dict_String metadata;
    Dict_Dict_String quotes;
} MarketData;

void release_MarketData(MarketData data);

DEFINE_LIST(MarketData)

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

DEFINE_RESULT(List_Category)


typedef struct {
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
   String last_updated;
} Quote;

void release_Quote(Quote data);

DEFINE_DICT(Quote)

typedef struct {
    Int id;
    String name;
    String symbol;
    String slug;
    Int num_market_pairs;
    String date_added;
    String tags;
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

DEFINE_RESULT(Optional_CategorySingle)
#endif
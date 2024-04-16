#ifndef EXCHANGE_H
#define EXCHANGE_H

#include <tools.h>

DEFINE_OPTIONAL(String)

DEFINE_OPTIONAL(Int)

DEFINE_OPTIONAL(Bool)

typedef struct {
    Int id;
    String name;
    String slug;
    Int is_active;
    Int is_listed;
    Int is_redistributable;
    String first_historical_data;
    String last_historical_data;
} Exchange;

void release_Exchange(Exchange data);

DEFINE_LIST(Exchange)

DEFINE_RESULT(List_Exchange)



DEFINE_LIST(String)

typedef struct {
    Int id;
    String name;
    String slug;
    String description;
    String notice;
    List_String fiats;
    String urls;
    String date_launched;
    Float maker_fee;
    Float taker_fee;
    Float spot_volume_usd;
    String spot_volume_last_updated;
    Int weekly_visits;
} Metadata;

void release_Metadata(Metadata data);

DEFINE_DICT(Metadata)

DEFINE_RESULT(Dict_Metadata)


typedef struct {
    Int crypto_id;
    String symbol;
    String name;
} Platform;

void release_Platform(Platform data);

typedef struct {
    Int crypto_id;
    Float price_usd;
    String symbol;
    String name;
} Currency;

void release_Currency(Currency data);

typedef struct {
    String wallet_address;
    Float balance;
    Platform platform;
    Currency currency;
} Asset;

void release_Asset(Asset data);

DEFINE_LIST(Asset)

DEFINE_RESULT(List_Asset)

#endif
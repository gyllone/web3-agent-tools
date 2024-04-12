#ifndef CRYPTOCURRENCY_H
#define CRYPTOCURRENCY_H

#include <tools.h>

DEFINE_OPTIONAL(String)

DEFINE_OPTIONAL(Int)

DEFINE_OPTIONAL(Bool)

DEFINE_LIST(Float)

DEFINE_LIST(List_Float)

DEFINE_DICT(String)

DEFINE_LIST(Dict_String)

DEFINE_DICT(Dict_String)

typedef struct {
    Bool is_fail;
    String error_message;
    List_List_Float data;
} QuoteResult;

void release_QuoteResult(QuoteResult result);

typedef struct {
    Bool is_fail;
    String error_message;
    List_Dict_String data;
} IdMapResult;

void release_IdMapResult(IdMapResult result);

typedef struct {
    Bool is_fail;
    String error_message;
    List_Dict_String data;
} MetadataResult;

void release_MetadataResult(MetadataResult result);

typedef struct {
    Dict_String metadata;
    Dict_Dict_String quotes;
} MarketData;

void release_MarketData(MarketData data);

DEFINE_LIST(MarketData)

typedef struct {
    Bool is_fail;
    String error_message;
    List_MarketData data;
}ListingResult;

void release_ListingResult(ListingResult result);

#endif
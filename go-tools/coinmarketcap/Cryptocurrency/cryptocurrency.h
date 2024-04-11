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

typedef struct {
    Bool is_fail;
    String error_message;
    List_List_Float quotes;
} QuoteResult;

void release_QuoteResult(QuoteResult result);

typedef struct {
    Bool is_fail;
    String error_message;
    List_Dict_String id_maps;
} IdMapResult;

void release_IdMapResult(IdMapResult result);

typedef struct {
    Bool is_fail;
    String error_message;
    List_Dict_String metas;
} MetadataResult;

void release_MetadataResult(MetadataResult result);


#endif
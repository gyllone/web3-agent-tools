#ifndef HASHKEY_H
#define HASHKEY_H

#include <tools.h>

DEFINE_OPTIONAL(String);
DEFINE_OPTIONAL(Int);

// balance
typedef struct {
    String Asset;
    String Total;
    String Free;
} Balance;

void release_Balance(Balance bal);

DEFINE_LIST(Balance)
DEFINE_RESULT(List_Balance)

// order
typedef struct {
    String OrderId;
    String SymbolName;
    String TransactTime;
    String Price;
    String Status;
    String OrigQty;
    String ExecutedQty;
} Order;

void release_Order(Order order);

DEFINE_LIST(Order)
DEFINE_RESULT(Order)
DEFINE_RESULT(List_Order)

// Kline
typedef struct {
    String Timestamp;
    String Symbol;
    String OpeningPrice;
    String ClosingPrice;
    String HighestPrice;
    String LowestPrice;
    String Volume;
} Kline;

void release_Kline(Kline kline);

DEFINE_LIST(Kline)
DEFINE_RESULT(List_Kline)

#endif
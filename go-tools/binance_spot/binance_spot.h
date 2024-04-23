#ifndef BINANCE_SPOT_H
#define BINANCE_SPOT_H

#include <tools.h>

DEFINE_OPTIONAL(String)
DEFINE_OPTIONAL(Int)
DEFINE_OPTIONAL(Float)
DEFINE_LIST(String)

// === Wallet ===

typedef struct {
    String id;
    Float amount;
    Float tx_fee;
    String coin;
    Int status;
    String address;
    String tx_id;
    String apply_time;
    String network;
    Int transfer_type;
    String info;
    Int confirmations;
    Int wallet_type;
} WithdrawHistory;

extern void release_WithdrawHistory(WithdrawHistory data);

DEFINE_LIST(WithdrawHistory)
DEFINE_OPTIONAL(List_WithdrawHistory)
DEFINE_RESULT(List_WithdrawHistory)

typedef struct {
    String asset;
    Float free;
    Float locked;
    Float freeze;
    Float withdrawing;
    Float btc_valuation;
} FundingAsset;

extern void release_FundingAsset(FundingAsset data);

DEFINE_LIST(FundingAsset)
DEFINE_OPTIONAL(List_FundingAsset)
DEFINE_RESULT(List_FundingAsset)

// === Market ===

typedef struct {
    String t;
    String s;
    Float o;
    Float c;
    Float h;
    Float l;
    Float v;
    Float q;
    Float tb;
    Float tq;
    Int n;
} KLine;
extern void release_KLine(KLine kline);

DEFINE_LIST(KLine)
DEFINE_OPTIONAL(List_KLine)
DEFINE_RESULT(List_KLine)

typedef struct {
    String symbol;
    Float price_change;
    Float price_change_pct;
    Float weighted_avg_price;
    Float last_price;
    Float open_price;
    Float high_price;
    Float low_price;
    Float volume;
    Float quote_volume;
    String open_time;
    String close_time;
    Int count;
} PriceChange;
extern void release_PriceChange(PriceChange pc);

DEFINE_LIST(PriceChange)
DEFINE_OPTIONAL(List_PriceChange)
DEFINE_RESULT(List_PriceChange)

// Get the latest price for symbols
typedef struct {
    String symbol;
    Float price;
} LatestPrice;
extern void release_LatestPrice(LatestPrice lp);

DEFINE_LIST(LatestPrice)
DEFINE_OPTIONAL(List_LatestPrice)
DEFINE_RESULT(List_LatestPrice)

// === Trade ===

typedef struct {
    String symbol;
    Int order_id;
    String transact_time;
    String working_time;
    Float price;
    Float orig_qty;
    Float executed_qty;
    Float cumulative_quote_qty;
    String status;
    String time_in_force;
    String order_type;
    String side;
} OrderResponse;
extern void release_OrderResponse(OrderResponse data);

DEFINE_OPTIONAL(OrderResponse)
DEFINE_RESULT(OrderResponse)

DEFINE_RESULT(Int)

typedef struct {
    String symbol;
    Int order_id;
    Float price;
    Float orig_qty;
    Float executed_qty;
    Float cumulative_quote_qty;
    String status;
    String time_in_force;
    String order_type;
    String side;
    Float stop_price;
    String timestamp;
    String update_time;
    Bool is_working;
    String working_time;
    Float orig_quote_order_qty;
} Order;
extern void release_Order(Order data);

DEFINE_LIST(Order)
DEFINE_OPTIONAL(List_Order)
DEFINE_RESULT(List_Order)

#endif
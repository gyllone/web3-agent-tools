#ifndef BINANCE_SPOT_H
#define BINANCE_SPOT_H

#include <tools.h>

DEFINE_OPTIONAL(String)
DEFINE_OPTIONAL(Int)
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
DEFINE_RESULT(List_FundingAsset)

// === Market ===

typedef struct {
    String open_time;
    Float open_price;
    Float high_price;
    Float low_price;
    Float close_price;
    Float volume;
    String close_time;
    Float quote_asset_volume;
    Int number_of_trades;
    Float taker_buy_base_asset_volume;
    Float taker_buy_quote_asset_volume;
} KLine;
extern void release_KLine(KLine kline);

DEFINE_LIST(KLine)
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
} PriceChange24h;
extern void release_PriceChange24h(PriceChange24h pc);

DEFINE_LIST(PriceChange24h)
DEFINE_RESULT(List_PriceChange24h)

// Get the rolling window price change for symbols
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
} RollingPriceChange;
extern void release_RollingPriceChange(RollingPriceChange rpc);

DEFINE_LIST(RollingPriceChange)
DEFINE_RESULT(List_RollingPriceChange)

// Get the latest price for symbols
typedef struct {
    String symbol;
    Float price;
} LatestPrice;
extern void release_LatestPrice(LatestPrice lp);

DEFINE_LIST(LatestPrice)
DEFINE_RESULT(List_LatestPrice)

#endif
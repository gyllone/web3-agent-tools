#include <binance_spot.h>

IMPL_OPTIONAL(String)
IMPL_OPTIONAL(Int)
IMPL_LIST(String)

// === Wallet ===

void release_WithdrawHistory(WithdrawHistory data) {
    free(data.id);
    free(data.coin);
    free(data.address);
    free(data.tx_id);
    free(data.apply_time);
    free(data.network);
    free(data.info);
}

IMPL_LIST(WithdrawHistory)
IMPL_RESULT(List_WithdrawHistory)

void release_FundingAsset(FundingAsset data) {
    free(data.asset);
}

IMPL_LIST(FundingAsset)
IMPL_RESULT(List_FundingAsset)

// === Market ===

void release_KLine(KLine kline) {
    free(kline.open_time);
    free(kline.close_time);
}

IMPL_LIST(KLine)
IMPL_RESULT(List_KLine)

void release_PriceChange24h(PriceChange24h pc) {
    free(pc.symbol);
    free(pc.open_time);
    free(pc.close_time);
}

IMPL_LIST(PriceChange24h)
IMPL_RESULT(List_PriceChange24h)

void release_RollingPriceChange(RollingPriceChange rpc) {
    free(rpc.symbol);
    free(rpc.open_time);
    free(rpc.close_time);
}

IMPL_LIST(RollingPriceChange)
IMPL_RESULT(List_RollingPriceChange)

void release_LatestPrice(LatestPrice lp) {
    free(lp.symbol);
}

IMPL_LIST(LatestPrice)
IMPL_RESULT(List_LatestPrice)

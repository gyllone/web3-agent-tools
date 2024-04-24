#include <binance_spot.h>

IMPL_OPTIONAL(String)
IMPL_OPTIONAL(Int)
IMPL_OPTIONAL(Float)
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
IMPL_OPTIONAL(List_WithdrawHistory)
IMPL_RESULT(List_WithdrawHistory)

void release_FundingAsset(FundingAsset data) {
    free(data.asset);
}

IMPL_LIST(FundingAsset)
IMPL_OPTIONAL(List_FundingAsset)
IMPL_RESULT(List_FundingAsset)

// === Market ===

void release_KLine(KLine kline) {
    free(kline.t);
    // kline.s is assigned in python, so we don't free it here
}

IMPL_LIST(KLine)
IMPL_OPTIONAL(List_KLine)
IMPL_RESULT(List_KLine)

void release_PriceChange(PriceChange pc) {
    free(pc.symbol);
    free(pc.open_time);
    free(pc.close_time);
}

IMPL_LIST(PriceChange)
IMPL_OPTIONAL(List_PriceChange)
IMPL_RESULT(List_PriceChange)

void release_LatestPrice(LatestPrice lp) {
    free(lp.symbol);
}

IMPL_LIST(LatestPrice)
IMPL_OPTIONAL(List_LatestPrice)
IMPL_RESULT(List_LatestPrice)

// === Trade ===

void release_OrderResponse(OrderResponse order) {
    free(order.symbol);
    free(order.transact_time);
    free(order.working_time);
    free(order.status);
    free(order.time_in_force);
    free(order.order_type);
    free(order.side);
}

IMPL_OPTIONAL(OrderResponse)
IMPL_RESULT(OrderResponse)

IMPL_RESULT(Int)

void release_Order(Order order) {
    free(order.symbol);
    free(order.status);
    free(order.time_in_force);
    free(order.order_type);
    free(order.side);
    free(order.timestamp);
    free(order.update_time);
    free(order.working_time);
}

IMPL_LIST(Order)
IMPL_OPTIONAL(List_Order)
IMPL_RESULT(List_Order)
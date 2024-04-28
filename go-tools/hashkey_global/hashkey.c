#include <hashkey.h>

IMPL_OPTIONAL(String);
IMPL_OPTIONAL(Int);

// balance
void release_Balance(Balance bal) {
    free(bal.Asset);
    free(bal.Total);
    free(bal.Free);
}
IMPL_LIST(Balance)
IMPL_OPTIONAL(List_Balance);
IMPL_RESULT(List_Balance)

// order
void release_Order(Order order) {
    free(order.OrderId);
    free(order.SymbolName);
    free(order.TransactTime);
    free(order.Price);
    free(order.Status);
    free(order.OrigQty);
    free(order.ExecutedQty);

}
IMPL_LIST(Order)
IMPL_OPTIONAL(Order)
IMPL_OPTIONAL(List_Order);
IMPL_RESULT(List_Order)
IMPL_RESULT(Order)

// kline
void release_Kline(Kline kline) {
    free(kline.Timestamp);
    free(kline.Symbol);
    free(kline.OpeningPrice);
    free(kline.ClosingPrice);
    free(kline.HighestPrice);
    free(kline.LowestPrice);
    free(kline.Volume);
}
IMPL_LIST(Kline)
IMPL_OPTIONAL(List_Kline);
IMPL_RESULT(List_Kline)

// price
void release_Price(Price price) {
    free(price.Symbol);
    free(price.Price);

}
IMPL_LIST(Price)
IMPL_OPTIONAL(List_Price);
IMPL_RESULT(List_Price)


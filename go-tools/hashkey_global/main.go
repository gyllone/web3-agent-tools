package main

/*
#cgo CFLAGS: -I../dependencies
#include <hashkey.h>
*/
import "C"
import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

func main() {
}

//export get_spot_trade_account_balance
func get_spot_trade_account_balance(apiKey C.String, secret C.String) C.Result_List_Balance {
	hashkeyAuth := __get_hashkey_auth(apiKey, secret)

	resp, err := getAccountInfo(&GetAccountInfoRequest{
		Timestamp: time.Now().UnixMilli(),
	}, hashkeyAuth)
	if err != nil {
		return C.err_List_Balance(C.CString(err.Error()))
	}

	bals := C.new_List_Balance(C.size_t(len(resp.Balances)))
	itemsSlice := (*[1 << 30]C.Balance)(unsafe.Pointer(bals.values))
	for i, bal := range resp.Balances {
		itemsSlice[i] = C.Balance{
			Asset: C.CString(bal.AssetName),
			Total: C.CString(bal.Total),
			Free:  C.CString(bal.Free),
		}
	}
	fmt.Printf("bals: %+v\n", bals)
	return C.ok_List_Balance(bals)
}

//export get_spot_trade_account_balance_release
func get_spot_trade_account_balance_release(res C.Result_List_Balance) {
	C.release_Result_List_Balance(res)
}

// create a spot order in market price
//
//export create_spot_market_order
func create_spot_market_order(apiKey C.String, secret C.String, symbol C.String, side C.String, quantity C.String) C.Result_Order {
	auth := __get_hashkey_auth(apiKey, secret)
	resp, err := createSpotOrder(&CreateSpotOrderRequest{
		Symbol:    C.GoString(symbol),
		Side:      SpotOrderSideEnum(C.GoString(side)),
		Type:      SpotOrderTypeEnum_MARKET,
		Quantity:  getPtr(C.GoString(quantity)),
		Timestamp: time.Now().UnixMilli(),
	}, auth)
	if err != nil {
		return C.err_Order(C.CString(err.Error()))
	}

	var order C.Order
	order.OrderId = C.CString(resp.OrderId)
	order.SymbolName = C.CString(resp.SymbolName)

	transacTimeInt64, e := strconv.ParseInt(resp.TransactTime, 10, 64)
	if e != nil {
		order.TransactTime = C.CString("")
	} else {
		order.TransactTime = C.CString(time.UnixMilli(transacTimeInt64).Format(time.RFC3339))
	}
	order.Price = C.CString(resp.Price)
	order.Status = C.CString(resp.Status)
	order.OrigQty = C.CString(resp.OrigQty)
	order.ExecutedQty = C.CString(resp.ExecutedQty)
	return C.ok_Order(order)

}

//export create_spot_market_order_release
func create_spot_market_order_release(res C.Result_Order) {
	C.release_Result_Order(res)
}

//export get_kline
func get_kline(Symbol C.String, Interval C.String, StartTime C.Optional_String, EndTime C.Optional_String, Limit C.Optional_String) C.Result_List_Kline {
	req := __get_kline_request(Symbol, Interval, StartTime, EndTime, Limit)
	resp, err := getQuoteKline(req)
	if err != nil {
		return C.err_List_Kline(C.CString(err.Error()))
	}

	klines := C.new_List_Kline(C.size_t(len(resp)))
	itemsSlice := (*[1 << 30]C.Kline)(unsafe.Pointer(klines.values))
	for i, item := range resp {
		itemsSlice[i] = C.Kline{
			Timestamp:    C.CString(time.UnixMilli(item.OpenTime).Format(time.RFC3339)),
			Symbol:       C.CString(req.Symbol),
			OpeningPrice: C.CString(item.OpenPrice),
			ClosingPrice: C.CString(item.ClosePrice),
			HighestPrice: C.CString(item.HighPrice),
			LowestPrice:  C.CString(item.LowPrice),
			Volume:       C.CString(item.Volume),
		}
	}
	return C.ok_List_Kline(klines)
}

//export get_kline_release
func get_kline_release(res C.Result_List_Kline) {
	C.release_Result_List_Kline(res)
}

// startTime:RFC3339
func __get_kline_request(Symbol C.String, Interval C.String, StartTime C.Optional_String, EndTime C.Optional_String, Limit C.Optional_String) *QuoteKlineRequest {
	return &QuoteKlineRequest{
		Symbol:    C.GoString(Symbol),
		Interval:  QuoteKlineENUM(C.GoString(Interval)),
		StartTime: __parse_string_to_i64(StartTime),
		EndTime:   __parse_string_to_i64(EndTime),
		Limit:     __parse_string_to_i(Limit),
	}
}

func __get_hashkey_auth(apiKey C.String, secret C.String) *HashKeyApiAuth {
	return &HashKeyApiAuth{
		ApiKey: C.GoString(apiKey),
		Secret: C.GoString(secret),
	}
}

func __parse_string_to_i64(cString C.Optional_String) *int64 {
	if !bool(cString.is_some) {
		return nil
	}
	if value, err := time.Parse(time.RFC3339, C.GoString(cString.value)); err != nil {
		return nil
	} else {
		return getPtr(value.UnixMilli())
	}
}

func __parse_string_to_i(cString C.Optional_String) *int {
	if !bool(cString.is_some) {
		return nil
	}
	goString := C.GoString(cString.value)
	if i, err := strconv.Atoi(goString); err != nil {
		return nil
	} else {
		return getPtr(i)
	}
}

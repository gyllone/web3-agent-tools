package main

/*
#cgo CFLAGS: -I../dependencies
#include <binance_spot.h>
*/
import "C"
import (
	"context"
	"encoding/json"
	"fmt"
	binance_connector "github.com/binance/binance-connector-go"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// === Wallet ===

////export withdraw_history
//func withdraw_history(
//	apiKey C.String,
//	secretKey C.String,
//	coin C.Optional_String,
//	withdrawId C.Optional_String,
//	status C.Optional_Int,
//	offset C.Optional_Int,
//	limit C.Optional_Int,
//	startTime C.Optional_Int,
//	endTime C.Optional_Int,
//) C.Result_List_WithdrawHistory {
//	cli := binance_connector.NewClient(C.GoString(apiKey), C.GoString(secretKey), BinanceUrl)
//	svc := cli.NewWithdrawHistoryService()
//	if bool(coin.is_some) {
//		svc.Coin(C.GoString(coin.value))
//	}
//	if bool(withdrawId.is_some) {
//		svc.WithdrawOrderId(C.GoString(withdrawId.value))
//	}
//	if bool(status.is_some) {
//		svc.Status(int(status.value))
//	}
//	if bool(offset.is_some) {
//		svc.Offset(int(offset.value))
//	}
//	if bool(limit.is_some) {
//		svc.Limit(int(limit.value))
//	}
//	if bool(startTime.is_some) {
//		svc.StartTime(uint64(startTime.value))
//	}
//	if bool(endTime.is_some) {
//		svc.EndTime(uint64(endTime.value))
//	}
//	resps, err := svc.Do(context.Background())
//	if err != nil {
//		return C.err_List_WithdrawHistory(C.CString(err.Error()))
//	}
//
//	data := C.new_List_WithdrawHistory(C.size_t(len(resps)))
//	dataSlice := (*[1 << 30]C.WithdrawHistory)(unsafe.Pointer(data.values))
//	for i, resp := range resps {
//		amount, _ := strconv.ParseFloat(resp.Amount, 64)
//		txFee, _ := strconv.ParseFloat(resp.TransactionFee, 64)
//		dataSlice[i] = C.WithdrawHistory{
//			id:            C.CString(resp.WithdrawOrderId),
//			amount:        C.Float(amount),
//			tx_fee:        C.Float(txFee),
//			coin:          C.CString(resp.Coin),
//			status:        C.Int(resp.Status),
//			address:       C.CString(resp.Address),
//			tx_id:         C.CString(resp.TxId),
//			apply_time:    C.CString(resp.ApplyTime),
//			network:       C.CString(resp.Network),
//			transfer_type: C.Int(resp.TransferType),
//			info:          C.CString(resp.Info),
//			confirmations: C.Int(resp.ConfirmNo),
//			wallet_type:   C.Int(resp.WalletType),
//		}
//	}
//	return C.ok_List_WithdrawHistory(data)
//}
//
////export withdraw_history_release
//func withdraw_history_release(output C.Result_List_WithdrawHistory) {
//	C.release_Result_List_WithdrawHistory(output)
//}

//export funding_asset
func funding_asset(
	apiKey C.String,
	secretKey C.String,
	asset C.Optional_String,
) C.Result_List_FundingAsset {
	cli := binance_connector.NewClient(C.GoString(apiKey), C.GoString(secretKey), BinanceUrl)
	svc := cli.NewFundingWalletService().NeedBtcValuation("true")
	if bool(asset.is_some) {
		svc.Asset(C.GoString(asset.value))
	}
	resps, err := svc.Do(context.Background())
	if err != nil {
		return C.err_List_FundingAsset(C.CString(err.Error()))
	}

	data := C.new_List_FundingAsset(C.size_t(len(resps)))
	dataSlice := (*[1 << 30]C.FundingAsset)(unsafe.Pointer(data.values))
	for i, resp := range resps {
		free, _ := strconv.ParseFloat(resp.Free, 64)
		locked, _ := strconv.ParseFloat(resp.Locked, 64)
		freeze, _ := strconv.ParseFloat(resp.Freeze, 64)
		withdrawing, _ := strconv.ParseFloat(resp.Withdrawing, 64)
		btcValuation, _ := strconv.ParseFloat(resp.BtcValuation, 64)
		dataSlice[i] = C.FundingAsset{
			asset:         C.CString(resp.Asset),
			free:          C.Float(free),
			locked:        C.Float(locked),
			freeze:        C.Float(freeze),
			withdrawing:   C.Float(withdrawing),
			btc_valuation: C.Float(btcValuation),
		}
	}
	return C.ok_List_FundingAsset(data)
}

//export funding_asset_release
func funding_asset_release(output C.Result_List_FundingAsset) {
	C.release_Result_List_FundingAsset(output)
}

//export get_account_info
func get_account_info(
	apiKey C.String,
	secretKey C.String,
	recvWindow C.Optional_Int,
) C.Result_AccountInfo {
	cli := binance_connector.NewClient(C.GoString(apiKey), C.GoString(secretKey), BinanceUrl)
	svc := cli.NewGetAccountService()
	var opts []binance_connector.RequestOption
	if bool(recvWindow.is_some) {
		opts = append(opts, binance_connector.WithRecvWindow(int64(recvWindow.value)))
	}

	resp, err := svc.Do(context.Background(), opts...)
	if err != nil {
		return C.err_AccountInfo(C.CString(err.Error()))
	}
	balances := C.new_List_Balance(C.size_t(len(resp.Balances)))
	balancesSlice := (*[1 << 30]C.Balance)(unsafe.Pointer(balances.values))
	for i, balance := range resp.Balances {
		free, _ := strconv.ParseFloat(balance.Free, 64)
		locked, _ := strconv.ParseFloat(balance.Locked, 64)
		balancesSlice[i] = C.Balance{
			asset:  C.CString(balance.Asset),
			free:   C.Float(free),
			locked: C.Float(locked),
		}
	}
	permissions := C.new_List_String(C.size_t(len(resp.Permissions)))
	permissionsSlice := (*[1 << 30]C.String)(unsafe.Pointer(permissions.values))
	for i, permission := range resp.Permissions {
		permissionsSlice[i] = C.CString(permission)
	}
	updateTime := time.UnixMilli(int64(resp.UpdateTime)).Format(time.RFC3339)
	accountInfo := C.AccountInfo{
		can_trade:    C.Bool(resp.CanTrade),
		can_withdraw: C.Bool(resp.CanWithdraw),
		can_deposit:  C.Bool(resp.CanDeposit),
		update_time:  C.CString(updateTime),
		balances:     balances,
		permissions:  permissions,
	}
	return C.ok_AccountInfo(accountInfo)
}

//export get_account_info_release
func get_account_info_release(output C.Result_AccountInfo) {
	C.release_Result_AccountInfo(output)
}

// === Market ===

//export k_lines
func k_lines(
	symbol C.String,
	interval C.String,
	limit C.Int,
	startTime C.Optional_String,
	endTime C.Optional_String,
) C.Result_List_KLine {
	cli := binance_connector.NewClient("", "", BinanceUrl)
	svc := cli.NewKlinesService().
		Symbol(C.GoString(symbol)).
		Interval(C.GoString(interval)).
		Limit(int(limit))
	if bool(startTime.is_some) {
		start, err := time.Parse(time.RFC3339, C.GoString(startTime.value))
		if err != nil {
			return C.err_List_KLine(C.CString(err.Error()))
		}
		svc.StartTime(uint64(start.UnixMilli()))
	}
	if bool(endTime.is_some) {
		end, err := time.Parse(time.RFC3339, C.GoString(endTime.value))
		if err != nil {
			return C.err_List_KLine(C.CString(err.Error()))
		}
		svc.EndTime(uint64(end.UnixMilli()))
	}

	resps, err := svc.Do(context.Background())
	if err != nil {
		return C.err_List_KLine(C.CString(err.Error()))
	}

	data := C.new_List_KLine(C.size_t(len(resps)))
	dataSlice := (*[1 << 30]C.KLine)(unsafe.Pointer(data.values))
	for i, resp := range resps {
		openTime := time.UnixMilli(int64(resp.OpenTime)).Format(time.RFC3339)
		openPrice, _ := strconv.ParseFloat(resp.Open, 64)
		closePrice, _ := strconv.ParseFloat(resp.Close, 64)
		highPrice, _ := strconv.ParseFloat(resp.High, 64)
		lowPrice, _ := strconv.ParseFloat(resp.Low, 64)
		volumn, _ := strconv.ParseFloat(resp.Volume, 64)
		quoteAssetVolume, _ := strconv.ParseFloat(resp.QuoteAssetVolume, 64)
		takerBuyBaseAssetVolume, _ := strconv.ParseFloat(resp.TakerBuyBaseAssetVolume, 64)
		takerBuyQuoteAssetVolume, _ := strconv.ParseFloat(resp.TakerBuyQuoteAssetVolume, 64)
		dataSlice[i] = C.KLine{
			t:  C.CString(openTime),
			s:  symbol,
			o:  C.Float(openPrice),
			c:  C.Float(closePrice),
			h:  C.Float(highPrice),
			l:  C.Float(lowPrice),
			v:  C.Float(volumn),
			q:  C.Float(quoteAssetVolume),
			tb: C.Float(takerBuyBaseAssetVolume),
			tq: C.Float(takerBuyQuoteAssetVolume),
			n:  C.Int(resp.NumberOfTrades),
		}
	}
	return C.ok_List_KLine(data)
}

//export k_lines_release
func k_lines_release(output C.Result_List_KLine) {
	C.release_Result_List_KLine(output)
}

//export price_change_24h_statistics
func price_change_24h_statistics(
	bases C.List_String,
	quote C.String,
	descending bool,
	limit C.Int,
) C.Result_List_PriceChange {
	goLimit := int(limit)
	goQuote := strings.ToUpper(C.GoString(quote))
	basesLen := int(bases.len)
	if basesLen > 100 {
		return C.err_List_PriceChange(C.CString("base tokens length exceeds 100"))
	}
	if goLimit > 100 {
		return C.err_List_PriceChange(C.CString("limit exceeds 100"))
	}

	basesSlice := (*[1 << 30]C.String)(unsafe.Pointer(bases.values))[:basesLen:basesLen]
	var goSymbols []string
	for _, base := range basesSlice {
		goBase := C.GoString(base)
		goSymbols = append(goSymbols, strings.ToUpper(goBase)+goQuote)
	}

	url := BinanceUrl + "/api/v3/ticker/24hr"
	if len(goSymbols) > 0 {
		goSymbolsStr, _ := json.Marshal(goSymbols)
		url += fmt.Sprintf("?symbols=%s", goSymbolsStr)
	}
	r, err := http.Get(url)
	if err != nil {
		return C.err_List_PriceChange(C.CString(err.Error()))
	}
	defer r.Body.Close()

	var responses []Ticker24hrResponse
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&responses); err != nil {
		return C.err_List_PriceChange(C.CString(err.Error()))
	}
	var filteredResponses []Ticker24hrResponse
	for _, resp := range responses {
		if strings.HasSuffix(resp.Symbol, goQuote) {
			filteredResponses = append(filteredResponses, resp)
		}
	}

	// sort by price change percent
	sort.SliceStable(filteredResponses, func(i, j int) bool {
		pct1, _ := strconv.ParseFloat(filteredResponses[i].PriceChangePercent, 64)
		pct2, _ := strconv.ParseFloat(filteredResponses[j].PriceChangePercent, 64)
		return descending == (pct1 > pct2)
	})
	if goLimit < len(filteredResponses) {
		filteredResponses = filteredResponses[:goLimit]
	}

	data := C.new_List_PriceChange(C.size_t(len(filteredResponses)))
	dataSlice := (*[1 << 30]C.PriceChange)(unsafe.Pointer(data.values))
	for i, resp := range filteredResponses {
		priceChange, _ := strconv.ParseFloat(resp.PriceChange, 64)
		priceChangePct, _ := strconv.ParseFloat(resp.PriceChangePercent, 64)
		weightedAvgPrice, _ := strconv.ParseFloat(resp.WeightedAvgPrice, 64)
		lastPrice, _ := strconv.ParseFloat(resp.LastPrice, 64)
		openPrice, _ := strconv.ParseFloat(resp.OpenPrice, 64)
		highPrice, _ := strconv.ParseFloat(resp.HighPrice, 64)
		lowPrice, _ := strconv.ParseFloat(resp.LowPrice, 64)
		volume, _ := strconv.ParseFloat(resp.Volume, 64)
		quoteVolume, _ := strconv.ParseFloat(resp.QuoteVolume, 64)
		openTime := time.UnixMilli(int64(resp.OpenTime)).Format(time.RFC3339)
		closeTime := time.UnixMilli(int64(resp.CloseTime)).Format(time.RFC3339)
		dataSlice[i] = C.PriceChange{
			symbol:             C.CString(resp.Symbol),
			price_change:       C.Float(priceChange),
			price_change_pct:   C.Float(priceChangePct),
			weighted_avg_price: C.Float(weightedAvgPrice),
			last_price:         C.Float(lastPrice),
			open_price:         C.Float(openPrice),
			high_price:         C.Float(highPrice),
			low_price:          C.Float(lowPrice),
			volume:             C.Float(volume),
			quote_volume:       C.Float(quoteVolume),
			open_time:          C.CString(openTime),
			close_time:         C.CString(closeTime),
			count:              C.Int(resp.Count),
		}
	}
	return C.ok_List_PriceChange(data)
}

//export price_change_24h_statistics_release
func price_change_24h_statistics_release(output C.Result_List_PriceChange) {
	C.release_Result_List_PriceChange(output)
}

//export rolling_window_price_change_statistics
func rolling_window_price_change_statistics(
	symbols C.List_String,
	descending bool,
	windowSize C.String,
) C.Result_List_PriceChange {
	symbolsLen := int(symbols.len)
	if symbolsLen == 0 || symbolsLen > 50 {
		return C.err_List_PriceChange(C.CString("symbols length must be between 1 and 50"))
	}

	symbolsSlice := (*[1 << 30]C.String)(unsafe.Pointer(symbols.values))[:symbolsLen:symbolsLen]
	var goSymbols []string
	for _, symbol := range symbolsSlice {
		goSymbols = append(goSymbols, C.GoString(symbol))
	}

	goSymbolsStr, _ := json.Marshal(goSymbols)
	u := fmt.Sprintf("%s/api/v3/ticker?symbols=%s&windowSize=%s", BinanceUrl, goSymbolsStr, C.GoString(windowSize))
	r, err := http.Get(u)
	if err != nil {
		return C.err_List_PriceChange(C.CString(err.Error()))
	}
	defer r.Body.Close()

	var responses []binance_connector.TickerResponse
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&responses); err != nil {
		return C.err_List_PriceChange(C.CString(err.Error()))
	}
	// sort by price change percent
	sort.SliceStable(responses, func(i, j int) bool {
		pct1, _ := strconv.ParseFloat(responses[i].PriceChangePercent, 64)
		pct2, _ := strconv.ParseFloat(responses[j].PriceChangePercent, 64)
		return descending == (pct1 > pct2)
	})

	data := C.new_List_PriceChange(C.size_t(len(responses)))
	dataSlice := (*[1 << 30]C.PriceChange)(unsafe.Pointer(data.values))
	for i, resp := range responses {
		priceChange, _ := strconv.ParseFloat(resp.PriceChange, 64)
		priceChangePct, _ := strconv.ParseFloat(resp.PriceChangePercent, 64)
		weightedAvgPrice, _ := strconv.ParseFloat(resp.WeightedAvgPrice, 64)
		lastPrice, _ := strconv.ParseFloat(resp.LastPrice, 64)
		openPrice, _ := strconv.ParseFloat(resp.OpenPrice, 64)
		highPrice, _ := strconv.ParseFloat(resp.HighPrice, 64)
		lowPrice, _ := strconv.ParseFloat(resp.LowPrice, 64)
		volume, _ := strconv.ParseFloat(resp.Volume, 64)
		quoteVolume, _ := strconv.ParseFloat(resp.QuoteVolume, 64)
		openTime := time.UnixMilli(int64(resp.OpenTime)).Format(time.RFC3339)
		closeTime := time.UnixMilli(int64(resp.CloseTime)).Format(time.RFC3339)
		dataSlice[i] = C.PriceChange{
			symbol:             C.CString(resp.Symbol),
			price_change:       C.Float(priceChange),
			price_change_pct:   C.Float(priceChangePct),
			weighted_avg_price: C.Float(weightedAvgPrice),
			last_price:         C.Float(lastPrice),
			open_price:         C.Float(openPrice),
			high_price:         C.Float(highPrice),
			low_price:          C.Float(lowPrice),
			volume:             C.Float(volume),
			quote_volume:       C.Float(quoteVolume),
			open_time:          C.CString(openTime),
			close_time:         C.CString(closeTime),
			count:              C.Int(resp.Count),
		}
	}
	return C.ok_List_PriceChange(data)
}

//export rolling_window_price_change_statistics_release
func rolling_window_price_change_statistics_release(output C.Result_List_PriceChange) {
	C.release_Result_List_PriceChange(output)
}

//export latest_price
func latest_price(
	symbols C.List_String,
	ascending bool,
	limit C.Int,
) C.Result_List_LatestPrice {
	goLimit := int(limit)
	symbolsLen := int(symbols.len)
	if symbolsLen > 100 {
		return C.err_List_LatestPrice(C.CString("symbols length exceeds 100"))
	}
	if int(limit) > 100 {
		return C.err_List_LatestPrice(C.CString("limit exceeds 100"))
	}

	symbolsSlice := (*[1 << 30]C.String)(unsafe.Pointer(symbols.values))[:symbolsLen:symbolsLen]
	var goSymbols []string
	for _, symbol := range symbolsSlice {
		goSymbols = append(goSymbols, C.GoString(symbol))
	}

	url := BinanceUrl + "/api/v3/ticker/price"
	if len(goSymbols) > 0 {
		goSymbolsStr, _ := json.Marshal(goSymbols)
		url += fmt.Sprintf("?symbols=%s", goSymbolsStr)
	}
	r, err := http.Get(url)
	if err != nil {
		return C.err_List_LatestPrice(C.CString(err.Error()))
	}
	defer r.Body.Close()

	var responses []binance_connector.TickerPriceResponse
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&responses); err != nil {
		return C.err_List_LatestPrice(C.CString(err.Error()))
	}
	// sort by price change percent
	sort.SliceStable(responses, func(i, j int) bool {
		price1, _ := strconv.ParseFloat(responses[i].Price, 64)
		price2, _ := strconv.ParseFloat(responses[j].Price, 64)
		return ascending == (price1 < price2)
	})
	if goLimit < len(responses) {
		responses = responses[:goLimit]
	}

	data := C.new_List_LatestPrice(C.size_t(len(responses)))
	dataSlice := (*[1 << 30]C.LatestPrice)(unsafe.Pointer(data.values))
	for i, resp := range responses {
		price, _ := strconv.ParseFloat(resp.Price, 64)
		dataSlice[i] = C.LatestPrice{
			symbol: C.CString(resp.Symbol),
			price:  C.Float(price),
		}
	}
	return C.ok_List_LatestPrice(data)
}

//export latest_price_release
func latest_price_release(output C.Result_List_LatestPrice) {
	C.release_Result_List_LatestPrice(output)
}

// === Trade ===

//export create_order
func create_order(
	apiKey C.String,
	secretKey C.String,
	symbol C.String,
	side C.String,
	orderType C.String,
	quantity C.Float,
	timeInForce C.Optional_String,
	price C.Optional_Float,
	stopPrice C.Optional_Float,
	recvWindow C.Optional_Int,
) C.Result_OrderResponse {
	cli := binance_connector.NewClient(C.GoString(apiKey), C.GoString(secretKey), BinanceUrl)
	svc := cli.NewCreateOrderService().
		Symbol(C.GoString(symbol)).
		Side(C.GoString(side)).
		Type(C.GoString(orderType)).
		Quantity(float64(quantity)).
		NewOrderRespType("RESULT")
	if bool(timeInForce.is_some) {
		svc.TimeInForce(C.GoString(timeInForce.value))
	}
	if bool(price.is_some) {
		svc.Price(float64(price.value))
	}
	if bool(stopPrice.is_some) {
		svc.StopPrice(float64(stopPrice.value))
	}
	var opts []binance_connector.RequestOption
	if bool(recvWindow.is_some) {
		opts = append(opts, binance_connector.WithRecvWindow(int64(recvWindow.value)))
	}

	resp, err := svc.Do(context.Background(), opts...)
	if err != nil {
		return C.err_OrderResponse(C.CString(err.Error()))
	}
	resultResp := resp.(*binance_connector.CreateOrderResponseRESULT)

	transcactTime := time.UnixMilli(int64(resultResp.TransactTime)).Format(time.RFC3339)
	workingTime := time.UnixMilli(int64(resultResp.WorkingTime)).Format(time.RFC3339)
	finalPrice, _ := strconv.ParseFloat(resultResp.Price, 64)
	origQty, _ := strconv.ParseFloat(resultResp.OrigQty, 64)
	executedQty, _ := strconv.ParseFloat(resultResp.ExecutedQty, 64)
	cumulativeQuoteQty, _ := strconv.ParseFloat(resultResp.CumulativeQuoteQty, 64)
	orderResp := C.OrderResponse{
		symbol:               C.CString(resultResp.Symbol),
		order_id:             C.Int(resultResp.OrderId),
		transact_time:        C.CString(transcactTime),
		working_time:         C.CString(workingTime),
		price:                C.Float(finalPrice),
		orig_qty:             C.Float(origQty),
		executed_qty:         C.Float(executedQty),
		cumulative_quote_qty: C.Float(cumulativeQuoteQty),
		status:               C.CString(resultResp.Status),
		time_in_force:        C.CString(resultResp.TimeInForce),
		order_type:           C.CString(resultResp.Type),
		side:                 C.CString(resultResp.Side),
	}
	return C.ok_OrderResponse(orderResp)
}

//export create_order_release
func create_order_release(output C.Result_OrderResponse) {
	C.release_Result_OrderResponse(output)
}

//export cancel_order
func cancel_order(
	apiKey C.String,
	secretKey C.String,
	symbol C.String,
	orderId C.Int,
	cancelRestrictions C.Optional_String,
	recvWindow C.Optional_Int,
) C.Result_Int {
	cli := binance_connector.NewClient(C.GoString(apiKey), C.GoString(secretKey), BinanceUrl)
	svc := cli.NewCancelOrderService().
		Symbol(C.GoString(symbol)).
		OrderId(int64(orderId))
	if bool(cancelRestrictions.is_some) {
		svc.CancelRestrictions(C.GoString(cancelRestrictions.value))
	}
	var opts []binance_connector.RequestOption
	if bool(recvWindow.is_some) {
		opts = append(opts, binance_connector.WithRecvWindow(int64(recvWindow.value)))
	}

	resp, err := svc.Do(context.Background(), opts...)
	if err != nil {
		return C.err_Int(C.CString(err.Error()))
	}
	return C.ok_Int(C.Int(resp.OrderId))
}

//export cancel_order_release
func cancel_order_release(output C.Result_Int) {
	C.release_Result_Int(output)
}

//export get_open_orders
func get_open_orders(
	apiKey C.String,
	secretKey C.String,
	symbol C.Optional_String,
	recvWindow C.Optional_Int,
) C.Result_List_Order {
	cli := binance_connector.NewClient(C.GoString(apiKey), C.GoString(secretKey), BinanceUrl)
	svc := cli.NewGetOpenOrdersService()
	if bool(symbol.is_some) {
		svc.Symbol(C.GoString(symbol.value))
	}
	var opts []binance_connector.RequestOption
	if bool(recvWindow.is_some) {
		opts = append(opts, binance_connector.WithRecvWindow(int64(recvWindow.value)))
	}

	resps, err := svc.Do(context.Background(), opts...)
	if err != nil {
		return C.err_List_Order(C.CString(err.Error()))
	}

	data := C.new_List_Order(C.size_t(len(resps)))
	dataSlice := (*[1 << 30]C.Order)(unsafe.Pointer(data.values))
	for i, resp := range resps {
		timestamp := time.UnixMilli(int64(resp.Time)).Format(time.RFC3339)
		updateTime := time.UnixMilli(int64(resp.UpdateTime)).Format(time.RFC3339)
		workingTime := time.UnixMilli(int64(resp.WorkingTime)).Format(time.RFC3339)
		finalPrice, _ := strconv.ParseFloat(resp.Price, 64)
		origQty, _ := strconv.ParseFloat(resp.OrigQty, 64)
		executedQty, _ := strconv.ParseFloat(resp.ExecutedQty, 64)
		cumulativeQuoteQty, _ := strconv.ParseFloat(resp.CumulativeQuoteQty, 64)
		stopPrice, _ := strconv.ParseFloat(resp.StopPrice, 64)
		origQuoteOrderQty, _ := strconv.ParseFloat(resp.OrigQuoteOrderQty, 64)
		dataSlice[i] = C.Order{
			symbol:               C.CString(resp.Symbol),
			order_id:             C.Int(resp.OrderId),
			price:                C.Float(finalPrice),
			orig_qty:             C.Float(origQty),
			executed_qty:         C.Float(executedQty),
			cumulative_quote_qty: C.Float(cumulativeQuoteQty),
			status:               C.CString(resp.Status),
			time_in_force:        C.CString(resp.TimeInForce),
			order_type:           C.CString(resp.Type),
			side:                 C.CString(resp.Side),
			stop_price:           C.Float(stopPrice),
			timestamp:            C.CString(timestamp),
			update_time:          C.CString(updateTime),
			is_working:           C.Bool(resp.IsWorking),
			working_time:         C.CString(workingTime),
			orig_quote_order_qty: C.Float(origQuoteOrderQty),
		}
	}
	return C.ok_List_Order(data)
}

//export get_open_orders_release
func get_open_orders_release(output C.Result_List_Order) {
	C.release_Result_List_Order(output)
}

func main() {}

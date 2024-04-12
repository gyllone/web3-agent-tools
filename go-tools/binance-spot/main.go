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
	"time"
	"unsafe"
)

// === Wallet ===

//export withdraw_history
func withdraw_history(
	apiKey C.String,
	secretKey C.String,
	coin C.Optional_String,
	withdrawId C.Optional_String,
	status C.Optional_Int,
	offset C.Optional_Int,
	limit C.Optional_Int,
	startTime C.Optional_Int,
	endTime C.Optional_Int,
) C.Result_List_WithdrawHistory {
	cli := binance_connector.NewClient(C.GoString(apiKey), C.GoString(secretKey))
	svc := cli.NewWithdrawHistoryService()
	if bool(coin.is_some) {
		svc.Coin(C.GoString(coin.value))
	}
	if bool(withdrawId.is_some) {
		svc.WithdrawOrderId(C.GoString(withdrawId.value))
	}
	if bool(status.is_some) {
		svc.Status(int(status.value))
	}
	if bool(offset.is_some) {
		svc.Offset(int(offset.value))
	}
	if bool(limit.is_some) {
		svc.Limit(int(limit.value))
	}
	if bool(startTime.is_some) {
		svc.StartTime(uint64(startTime.value))
	}
	if bool(endTime.is_some) {
		svc.EndTime(uint64(endTime.value))
	}
	resps, err := svc.Do(context.Background())
	if err != nil {
		return C.err_List_WithdrawHistory(C.CString(err.Error()))
	}

	data := C.new_List_WithdrawHistory(C.size_t(len(resps)))
	dataSlice := (*[1 << 30]C.WithdrawHistory)(unsafe.Pointer(data.values))
	for i, resp := range resps {
		amount, _ := strconv.ParseFloat(resp.Amount, 64)
		txFee, _ := strconv.ParseFloat(resp.TransactionFee, 64)
		dataSlice[i] = C.WithdrawHistory{
			id:            C.CString(resp.WithdrawOrderId),
			amount:        C.Float(amount),
			tx_fee:        C.Float(txFee),
			coin:          C.CString(resp.Coin),
			status:        C.Int(resp.Status),
			address:       C.CString(resp.Address),
			tx_id:         C.CString(resp.TxId),
			apply_time:    C.CString(resp.ApplyTime),
			network:       C.CString(resp.Network),
			transfer_type: C.Int(resp.TransferType),
			info:          C.CString(resp.Info),
			confirmations: C.Int(resp.ConfirmNo),
			wallet_type:   C.Int(resp.WalletType),
		}
	}
	return C.ok_List_WithdrawHistory(data)
}

//export withdraw_history_release
func withdraw_history_release(output C.Result_List_WithdrawHistory) {
	C.release_Result_List_WithdrawHistory(output)
}

//export funding_asset
func funding_asset(
	apiKey C.String,
	secretKey C.String,
	asset C.Optional_String,
) C.Result_List_FundingAsset {
	cli := binance_connector.NewClient(C.GoString(apiKey), C.GoString(secretKey))
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

// === Market ===

//export k_lines
func k_lines(
	symbol C.String,
	interval C.String,
	limit C.Int,
	startTime C.Optional_Int,
	endTime C.Optional_Int,
) C.Result_List_KLine {
	cli := binance_connector.NewClient("", "")
	svc := cli.NewKlinesService().
		Symbol(C.GoString(symbol)).
		Interval(C.GoString(interval)).
		Limit(int(limit))
	if bool(startTime.is_some) {
		svc.StartTime(uint64(startTime.value))
	}
	if bool(endTime.is_some) {
		svc.EndTime(uint64(endTime.value))
	}

	resps, err := svc.Do(context.Background(), binance_connector.WithRecvWindow(RecvWindow))
	if err != nil {
		return C.err_List_KLine(C.CString(err.Error()))
	}

	data := C.new_List_KLine(C.size_t(len(resps)))
	dataSlice := (*[1 << 30]C.KLine)(unsafe.Pointer(data.values))
	for i, resp := range resps {
		openTime := time.UnixMilli(int64(resp.OpenTime)).Format(time.RFC3339)
		openPrice, _ := strconv.ParseFloat(resp.Open, 64)
		highPrice, _ := strconv.ParseFloat(resp.High, 64)
		lowPrice, _ := strconv.ParseFloat(resp.Low, 64)
		closePrice, _ := strconv.ParseFloat(resp.Close, 64)
		volume, _ := strconv.ParseFloat(resp.Volume, 64)
		closeTime := time.UnixMilli(int64(resp.CloseTime)).Format(time.RFC3339)
		quoteVolumn, _ := strconv.ParseFloat(resp.QuoteAssetVolume, 64)
		takerBuyBaseAssetVolume, _ := strconv.ParseFloat(resp.TakerBuyBaseAssetVolume, 64)
		takerBuyQuoteAssetVolume, _ := strconv.ParseFloat(resp.TakerBuyQuoteAssetVolume, 64)
		dataSlice[i] = C.KLine{
			open_time:                    C.CString(openTime),
			open_price:                   C.Float(openPrice),
			high_price:                   C.Float(highPrice),
			low_price:                    C.Float(lowPrice),
			close_price:                  C.Float(closePrice),
			volume:                       C.Float(volume),
			close_time:                   C.CString(closeTime),
			quote_asset_volume:           C.Float(quoteVolumn),
			number_of_trades:             C.Int(resp.NumberOfTrades),
			taker_buy_base_asset_volume:  C.Float(takerBuyBaseAssetVolume),
			taker_buy_quote_asset_volume: C.Float(takerBuyQuoteAssetVolume),
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
	symbols C.List_String,
	ascending bool,
	limit C.Int,
) C.Result_List_PriceChange24h {
	symbolsLen := int(symbols.len)
	if symbolsLen > 20 {
		return C.err_List_PriceChange24h(C.CString("symbols length exceeds 20"))
	}
	if int(limit) > 20 {
		return C.err_List_PriceChange24h(C.CString("limit exceeds 20"))
	}

	symbolsSlice := (*[1 << 30]C.String)(unsafe.Pointer(symbols.values))[:symbolsLen:symbolsLen]
	var goSymbols []string
	for _, symbol := range symbolsSlice {
		goSymbols = append(goSymbols, C.GoString(symbol))
	}

	url := BinanceUrl + "/api/v3/ticker/24hr"
	if len(goSymbols) > 0 {
		goSymbolsStr, _ := json.Marshal(goSymbols)
		url += fmt.Sprintf("?symbols=%s", goSymbolsStr)
	}
	r, err := http.Get(url)
	if err != nil {
		return C.err_List_PriceChange24h(C.CString(err.Error()))
	}
	defer r.Body.Close()

	var responses []binance_connector.Ticker24hrResponse
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&responses); err != nil {
		return C.err_List_PriceChange24h(C.CString(err.Error()))
	}

	// sort by price change percent
	sort.SliceStable(&responses, func(i, j int) bool {
		pct1, _ := strconv.ParseFloat(responses[i].PriceChangePercent, 64)
		pct2, _ := strconv.ParseFloat(responses[j].PriceChangePercent, 64)
		return ascending == (pct1 < pct2)
	})
	responses = responses[:limit]
	data := C.new_List_PriceChange24h(C.size_t(len(responses)))
	dataSlice := (*[1 << 30]C.PriceChange24h)(unsafe.Pointer(data.values))
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
		dataSlice[i] = C.PriceChange24h{
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
	return C.ok_List_PriceChange24h(data)
}

//export price_change_24h_statistics_release
func price_change_24h_statistics_release(output C.Result_List_PriceChange24h) {
	C.release_Result_List_PriceChange24h(output)
}

//export rolling_window_price_change_statistics
func rolling_window_price_change_statistics(
	symbols C.List_String,
	windowSize C.String,
) C.Result_List_RollingPriceChange {
	symbolsLen := int(symbols.len)
	if symbolsLen == 0 || symbolsLen > 50 {
		return C.err_List_RollingPriceChange(C.CString("symbols length must be between 1 and 50"))
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
		return C.err_List_RollingPriceChange(C.CString(err.Error()))
	}
	defer r.Body.Close()

	var responses []binance_connector.TickerResponse
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&responses); err != nil {
		return C.err_List_RollingPriceChange(C.CString(err.Error()))
	}

	data := C.new_List_RollingPriceChange(C.size_t(len(responses)))
	dataSlice := (*[1 << 30]C.RollingPriceChange)(unsafe.Pointer(data.values))
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
		dataSlice[i] = C.RollingPriceChange{
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
	return C.ok_List_RollingPriceChange(data)
}

//export rolling_window_price_change_statistics_release
func rolling_window_price_change_statistics_release(output C.Result_List_RollingPriceChange) {
	C.release_Result_List_RollingPriceChange(output)
}

//export latest_price
func latest_price(
	symbols C.List_String,
	ascending bool,
	limit C.Int,
) C.Result_List_LatestPrice {
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
	sort.SliceStable(&responses, func(i, j int) bool {
		price1, _ := strconv.ParseFloat(responses[i].Price, 64)
		price2, _ := strconv.ParseFloat(responses[j].Price, 64)
		return ascending == (price1 < price2)
	})
	responses = responses[:limit]
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

func main() {}

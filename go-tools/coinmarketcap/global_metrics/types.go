package main

import "time"

type StatusResp struct {
	ErrorMessage string `json:"error_message"`
}

type QuotesLatestResp struct {
	Data   Metric     `json:"data"`
	Status StatusResp `json:"status"`
}

type Metric struct {
	ActiveCryptocurrencies          int64            `json:"active_cryptocurrencies"`
	TotalCryptocurrencies           int64            `json:"total_cryptocurrencies"`
	ActiveMarketPairs               int64            `json:"active_market_pairs"`
	ActiveExchanges                 int64            `json:"active_exchanges"`
	TotalExchanges                  int64            `json:"total_exchanges"`
	EthDominance                    float64          `json:"eth_dominance"`
	BtcDominance                    float64          `json:"btc_dominance"`
	EthDominanceYesterday           float64          `json:"eth_dominance_yesterday"`
	BtcDominanceYesterday           float64          `json:"btc_dominance_yesterday"`
	EthDominance24hPercentageChange float64          `json:"eth_dominance_24h_percentage_change"`
	BtcDominance24hPercentageChange float64          `json:"btc_dominance_24h_percentage_change"`
	DefiVolume24h                   float64          `json:"defi_volume_24h"`
	DefiVolume24hReported           float64          `json:"defi_volume_24h_reported"`
	DefiMarketCap                   float64          `json:"defi_market_cap"`
	Defi24hPercentageChange         float64          `json:"defi_24h_percentage_change"`
	StablecoinVolume24h             float64          `json:"stablecoin_volume_24h"`
	StablecoinVolume24hReported     float64          `json:"stablecoin_volume_24h_reported"`
	StablecoinMarketCap             float64          `json:"stablecoin_market_cap"`
	Stablecoin24hPercentageChange   float64          `json:"stablecoin_24h_percentage_change"`
	DerivativesVolume24h            float64          `json:"derivatives_volume_24h"`
	DerivativesVolume24hReported    float64          `json:"derivatives_volume_24h_reported"`
	Derivatives24hPercentageChange  float64          `json:"derivatives_24h_percentage_change"`
	Quote                           map[string]Quote `json:"quote"`
	LastUpdated                     time.Time        `json:"last_updated"`
}

type Quote struct {
	TotalMarketCap                          float64   `json:"total_market_cap"`
	TotalVolume24h                          float64   `json:"total_volume_24h"`
	TotalVolume24hReported                  float64   `json:"total_volume_24h_reported"`
	AltcoinVolume24h                        float64   `json:"altcoin_volume_24h"`
	AltcoinVolume24hReported                float64   `json:"altcoin_volume_24h_reported"`
	AltcoinMarketCap                        float64   `json:"altcoin_market_cap"`
	DefiVolume24h                           float64   `json:"defi_volume_24h"`
	DefiVolume24hReported                   float64   `json:"defi_volume_24h_reported"`
	Defi24hPercentageChange                 float64   `json:"defi_24h_percentage_change"`
	DefiMarketCap                           float64   `json:"defi_market_cap"`
	StablecoinVolume24h                     float64   `json:"stablecoin_volume_24h"`
	StablecoinVolume24hReported             float64   `json:"stablecoin_volume_24h_reported"`
	Stablecoin24hPercentageChange           float64   `json:"stablecoin_24h_percentage_change"`
	StablecoinMarketCap                     float64   `json:"stablecoin_market_cap"`
	DerivativesVolume24h                    float64   `json:"derivatives_volume_24h"`
	DerivativesVolume24hReported            float64   `json:"derivatives_volume_24h_reported"`
	Derivatives24hPercentageChange          float64   `json:"derivatives_24h_percentage_change"`
	TotalMarketCapYesterday                 float64   `json:"total_market_cap_yesterday"`
	TotalVolume24hYesterday                 float64   `json:"total_volume_24h_yesterday"`
	TotalMarketCapYesterdayPercentageChange float64   `json:"total_market_cap_yesterday_percentage_change"`
	TotalVolume24hYesterdayPercentageChange float64   `json:"total_volume_24h_yesterday_percentage_change"`
	LastUpdated                             time.Time `json:"last_updated"`
}

package main

import "time"

type StatusResp struct {
	ErrorMessage string `json:"error_message"`
}

type Exchange struct {
	ID                  int64     `json:"id"`
	Name                string    `json:"name"`
	Slug                string    `json:"slug"`
	IsActive            int64     `json:"is_active"`
	IsListed            int64     `json:"is_listed"`
	IsRedistributable   int64     `json:"is_redistributable"`
	FirstHistoricalData time.Time `json:"first_historical_data"`
	LastHistoricalData  time.Time `json:"last_historical_data"`
}

type IdMapResp struct {
	Data   []Exchange `json:"data"`
	Status StatusResp `json:"status"`
}

type Metadata struct {
	ID                    int64     `json:"id"`
	Name                  string    `json:"name"`
	Slug                  string    `json:"slug"`
	Description           string    `json:"description"`
	Notice                string    `json:"notice"`
	Fiats                 []string  `json:"fiats"`
	URLs                  URLs      `json:"urls"`
	DateLaunched          time.Time `json:"date_launched"`
	MakerFee              float64   `json:"maker_fee"`
	TakerFee              float64   `json:"taker_fee"`
	SpotVolumeUSD         float64   `json:"spot_volume_usd"`
	SpotVolumeLastUpdated time.Time `json:"spot_volume_last_updated"`
	WeeklyVisits          int64     `json:"weekly_visits"`
}

type URLs struct {
	Twitter []string `json:"twitter"`
	Blog    []string `json:"blog"`
	Website []string `json:"website"`
	Chat    []string `json:"chat"`
	Actual  []string `json:"actual"`
	Fee     []string `json:"fee"`
}

type MetadataResp struct {
	Data   map[string]Metadata `json:"data"`
	Status StatusResp          `json:"status"`
}

type Asset struct {
	WalletAddress string  `json:"wallet_address"`
	Balance       float64 `json:"balance"`
	Platform      struct {
		CryptoID int64  `json:"crypto_id"`
		Symbol   string `json:"symbol"`
		Name     string `json:"name"`
	} `json:"platform"`
	Currency struct {
		CryptoID int64   `json:"crypto_id"`
		PriceUSD float64 `json:"price_usd"`
		Symbol   string  `json:"symbol"`
		Name     string  `json:"name"`
	} `json:"currency"`
}

type AssetResp struct {
	Data   []Asset    `json:"data"`
	Status StatusResp `json:"status"`
}

package main

import (
	"strconv"
	"strings"
	"time"
)

type StatusResp struct {
	ErrorMessage string `json:"error_message"`
}

type QuoteData struct {
	ID                            int64            `json:"id"`
	Name                          string           `json:"name"`
	Symbol                        string           `json:"symbol"`
	Slug                          string           `json:"slug"`
	NumMarketPairs                int64            `json:"num_market_pairs"`
	DateAdded                     time.Time        `json:"date_added"`
	Tags                          []Tags           `json:"tags"`
	MaxSupply                     int64            `json:"max_supply"`
	CirculatingSupply             float64          `json:"circulating_supply"`
	TotalSupply                   float64          `json:"total_supply"`
	IsActive                      int64            `json:"is_active"`
	InfiniteSupply                bool             `json:"infinite_supply"`
	Platform                      *Platform        `json:"platform"`
	CmcRank                       int64            `json:"cmc_rank"`
	IsFiat                        int64            `json:"is_fiat"`
	SelfReportedCirculatingSupply int64            `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         float64          `json:"self_reported_market_cap"`
	TvlRatio                      int64            `json:"tvl_ratio"`
	LastUpdated                   time.Time        `json:"last_updated"`
	Quote                         map[string]Quote `json:"quote"`
}

type Tags struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Quote struct {
	Price                 float64   `json:"price"`
	Volume24h             float64   `json:"volume_24h"`
	VolumeChange24h       float64   `json:"volume_change_24h"`
	PercentChange1h       float64   `json:"percent_change_1h"`
	PercentChange24h      float64   `json:"percent_change_24h"`
	PercentChange7d       float64   `json:"percent_change_7d"`
	PercentChange30d      float64   `json:"percent_change_30d"`
	PercentChange60d      float64   `json:"percent_change_60d"`
	PercentChange90d      float64   `json:"percent_change_90d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	Tvl                   float64   `json:"tvl"`
	LastUpdated           time.Time `json:"last_updated"`
}

type platformIdType int64

func (t *platformIdType) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*t = platformIdType(v)
	return nil
}

type Platform struct {
	ID           platformIdType `json:"id"`
	Name         string         `json:"name"`
	Symbol       string         `json:"symbol"`
	Slug         string         `json:"slug"`
	TokenAddress string         `json:"token_address,contract_address"`
}

type QuoteResp struct {
	Data   map[string]QuoteData `json:"data"`
	Status StatusResp           `json:"status"`
}

type Cryptocurrency struct {
	ID                  int64     `json:"id"`
	Rank                int64     `json:"rank"`
	Name                string    `json:"name"`
	Symbol              string    `json:"symbol"`
	Slug                string    `json:"slug"`
	IsActive            int64     `json:"is_active"`
	FirstHistoricalData time.Time `json:"first_historical_data"`
	LastHistoricalData  time.Time `json:"last_historical_data"`
	Platform            *Platform `json:"platform"`
}

type IdMapResp struct {
	Data   []Cryptocurrency `json:"data"`
	Status StatusResp       `json:"status"`
}

type Metadata struct {
	ID                            int64     `json:"id"`
	Name                          string    `json:"name"`
	Symbol                        string    `json:"symbol"`
	Category                      string    `json:"category"`
	Description                   string    `json:"description"`
	Slug                          string    `json:"slug"`
	Logo                          string    `json:"logo"`
	Subreddit                     string    `json:"subreddit"`
	Notice                        string    `json:"notice"`
	Tags                          []string  `json:"tags"`
	TagNames                      []string  `json:"tag_names"`
	TagGroups                     []string  `json:"tag_groups"`
	URLs                          URLs      `json:"urls"`
	Platform                      *Platform `json:"platform"`
	DateAdded                     time.Time `json:"date_added"`
	TwitterUsername               string    `json:"twitter_username"`
	IsHidden                      int64     `json:"is_hidden"`
	DateLaunched                  time.Time `json:"date_launched"`
	SelfReportedCirculatingSupply float64   `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         float64   `json:"self_reported_market_cap"`
	InfiniteSupply                bool      `json:"infinite_supply"`
}

type URLs struct {
	Website      []string `json:"website"`
	TechnicalDoc []string `json:"technical_doc"`
	Explorer     []string `json:"explorer"`
	SourceCode   []string `json:"source_code"`
	MessageBoard []string `json:"message_board"`
	Chat         []string `json:"chat"`
	Facebook     []string `json:"facebook"`
	Announcement []string `json:"announcement"`
	Reddit       []string `json:"reddit"`
	Twitter      []string `json:"twitter"`
}

type MetadataResp struct {
	Data   map[string]Metadata `json:"data"`
	Status StatusResp          `json:"status"`
}

type ListingsData struct {
	ID                            int64            `json:"id"`
	Name                          string           `json:"name"`
	Symbol                        string           `json:"symbol"`
	Slug                          string           `json:"slug"`
	CmcRank                       int64            `json:"cmc_rank"`
	NumMarketPairs                int64            `json:"num_market_pairs"`
	CirculatingSupply             float64          `json:"circulating_supply"`
	TotalSupply                   float64          `json:"total_supply"`
	MaxSupply                     float64          `json:"max_supply"`
	InfiniteSupply                bool             `json:"infinite_supply"`
	LastUpdated                   time.Time        `json:"last_updated"`
	DateAdded                     time.Time        `json:"date_added"`
	SelfReportedCirculatingSupply float64          `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         float64          `json:"self_reported_market_cap"`
	TvlRatio                      float64          `json:"tvl_ratio"`
	Tags                          []string         `json:"tags"`
	Platform                      *Platform        `json:"platform"`
	Quote                         map[string]Quote `json:"quote"`
}

type ListingsResp struct {
	Data   []ListingsData `json:"data"`
	Status StatusResp     `json:"status"`
}

type Category struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	NumTokens       int64     `json:"num_tokens"`
	AvgPriceChange  float64   `json:"avg_price_change"`
	MarketCap       float64   `json:"market_cap"`
	MarketCapChange float64   `json:"market_cap_change"`
	Volume          float64   `json:"volume"`
	VolumeChange    float64   `json:"volume_change"`
	LastUpdated     time.Time `json:"last_updated"`
}

type CategoriesResp struct {
	Data   []Category `json:"data"`
	Status StatusResp `json:"status"`
}

type CategorySingle struct {
	Category
	Coins []Coin `json:"coins"`
}

type Coin struct {
	ID                int64            `json:"id"`
	Name              string           `json:"name"`
	Symbol            string           `json:"symbol"`
	Slug              string           `json:"slug"`
	NumMarketPairs    int64            `json:"num_market_pairs"`
	DateAdded         time.Time        `json:"date_added"`
	Tags              []string         `json:"tags"`
	MaxSupply         float64          `json:"max_supply"`
	CirculatingSupply float64          `json:"circulating_supply"`
	TotalSupply       float64          `json:"total_supply"`
	IsActive          int64            `json:"is_active"`
	InfiniteSupply    bool             `json:"infinite_supply"`
	TvlRatio          float64          `json:"tvl_ratio"`
	CmcRank           int64            `json:"cmc_rank"`
	IsFiat            int64            `json:"is_fiat"`
	LastUpdated       time.Time        `json:"last_updated"`
	Quote             map[string]Quote `json:"quote"`
}

type CategoryResp struct {
	Data   CategorySingle `json:"data"`
	Status StatusResp     `json:"status"`
}

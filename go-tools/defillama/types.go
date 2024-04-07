package main

type ProtocolTvl struct {
	Name     string             `json:"name"`
	TVL      float64            `json:"tvl"`
	ChainTvl map[string]float64 `json:"chainTvls"`
}

type ChainTVL struct {
	Name string  `json:"name"`
	TVL  float64 `json:"tvl"`
}

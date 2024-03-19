package main

type ProtocolTvl struct {
	Name     string             `json:"name"`
	Chains   []string           `json:"chains"`
	ChainTvl map[string]float64 `json:"chainTvls"`
}

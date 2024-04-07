package main

import "C"
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const BaseUrl = "https://api.llama.fi"

func parseResultWithNameWithBlockchain(name, blockchain string) (float64, error) {
	resp, err := http.Get(BaseUrl + "/protocols")
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("response code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var tvls []ProtocolTvl
	if err = decoder.Decode(&tvls); err != nil {
		return 0, err
	}

	for _, info := range tvls {
		if strings.ToLower(info.Name) == strings.ToLower(name) {
			for chain, tvl := range info.ChainTvl {
				if strings.ToLower(chain) == strings.ToLower(blockchain) {
					return tvl, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("name %s or blockchain %s not found", name, blockchain)
}

func parseResultWithoutNameWithBlockchain(blockchain string) (float64, error) {
	resp, err := http.Get(BaseUrl + "/v2/chains")
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("response code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var tvls []ChainTVL
	if err = decoder.Decode(&tvls); err != nil {
		return 0, err
	}

	for _, info := range tvls {
		if strings.ToLower(info.Name) == strings.ToLower(blockchain) {
			return info.TVL, nil
		}
	}
	return 0, fmt.Errorf("blockchain %s not found", blockchain)
}

func parseResultNameWithoutBlockchain(name string) (float64, error) {
	resp, err := http.Get(BaseUrl + "/tvl/" + name)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("response code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return strconv.ParseFloat(string(data), 64)
}

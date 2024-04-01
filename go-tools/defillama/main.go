package main

import "C"
import (
	"encoding/json"
	"net/http"
	"strings"
)

/*
#include <stdbool.h>

typedef struct {
	bool is_some;
	char* value;
} OptionalStr;

typedef struct {
	double tvl;
} TVLResult;
*/
import "C"

//export query_tvl
func query_tvl(protocol *C.char, blockchain *C.char) C.TVLResult {
	//func query_tvl(protocol C.OptionalStr, blockchain C.OptionalStr) C.TVLResult {
	//nameIsSome := bool(name.is_some)
	//blockchainIsSome := bool(blockchain.is_some)
	//goProtocol := C.GoString(protocol.value)
	//goBlockchain := C.GoString(blockchain.value)
	goProtocol := C.GoString(protocol)
	goBlockchain := C.GoString(blockchain)

	resp, err := http.Get("https://api.llama.fi/protocols")
	if err != nil {
		return C.TVLResult{}
	}
	if resp.StatusCode != 200 {
		return C.TVLResult{}
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var tvls []ProtocolTvl
	if err = decoder.Decode(&tvls); err != nil {
		return C.TVLResult{}
	}
	for _, info := range tvls {
		if strings.ToLower(info.Name) == strings.ToLower(goProtocol) {
			for chain, tvl := range info.ChainTvl {
				if strings.ToLower(chain) == strings.ToLower(goBlockchain) {
					return C.TVLResult{tvl: C.double(tvl)}
				}
			}
		}
	}
	return C.TVLResult{}
}

func main() {}

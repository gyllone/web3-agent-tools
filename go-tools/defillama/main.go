package main

import "C"

/*
#include <stdlib.h>
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
import (
	"encoding/json"
	"net/http"
	"strings"
	"unsafe"
)

//export query_tvl
func query_tvl(protocol *C.OptionalStr, blockchain *C.OptionalStr) *C.TVLResult {
	//nameIsSome := bool(name.is_some)
	//blockchainIsSome := bool(blockchain.is_some)
	goProtocol := C.GoString(protocol.value)
	goBlockchain := C.GoString(blockchain.value)

	resp, err := http.Get("https://api.llama.fi/protocols")
	if err != nil {
		return nil
	}
	if resp.StatusCode != 200 {
		return nil
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var tvls []ProtocolTvl
	if err = decoder.Decode(&tvls); err != nil {
		return nil
	}
	for _, info := range tvls {
		if strings.ToLower(info.Name) == strings.ToLower(goProtocol) {
			for chain, tvl := range info.ChainTvl {
				if strings.ToLower(chain) == strings.ToLower(goBlockchain) {
					// malloc for result
					result := (*C.TVLResult)(C.malloc(C.sizeof_TVLResult))
					if result == nil {
						return nil
					} else {
						result.tvl = C.double(tvl)
						return result
					}
				}
			}
		}
	}
	return nil
}

//export query_tvl_release
func query_tvl_release(res *C.TVLResult) {
	C.free(unsafe.Pointer(res))
}

func main() {}

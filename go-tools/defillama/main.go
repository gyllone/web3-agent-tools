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
	bool status;
	char* error;
	double tvl;
} TVLResult;
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export query_tvl
func query_tvl(cName *C.OptionalStr, cBlockchain *C.OptionalStr) *C.TVLResult {
	// malloc for result
	result := (*C.TVLResult)(C.malloc(C.sizeof_TVLResult))
	if result == nil {
		return nil
	}

	if cName == nil || cBlockchain == nil {
		result.status = false
		result.error = C.CString("name or blockchain is nil pointer")
		return result
	}

	nameIsSome := bool(cName.is_some)
	blockchainIsSome := bool(cBlockchain.is_some)
	goName := C.GoString(cName.value)
	goBlockchain := C.GoString(cBlockchain.value)

	var tvl float64
	var err error
	if nameIsSome && blockchainIsSome {
		tvl, err = parseResultWithNameWithBlockchain(goName, goBlockchain)
	} else if !nameIsSome && blockchainIsSome {
		tvl, err = parseResultWithoutNameWithBlockchain(goBlockchain)
	} else if nameIsSome && !blockchainIsSome {
		tvl, err = parseResultNameWithoutBlockchain(goName)
	} else {
		err = fmt.Errorf("name and blockchain are not provided")
	}

	if err != nil {
		result.status = false
		result.error = C.CString(err.Error())
		return result
	}
	result.status = true
	result.error = C.CString("")
	result.tvl = C.double(tvl)
	return result
}

//export query_tvl_release
func query_tvl_release(res *C.TVLResult) {
	if res == nil {
		return
	}
	if res.error != nil {
		C.free(unsafe.Pointer(res.error))
	}
	C.free(unsafe.Pointer(res))
}

func main() {}

package main

/*
#include <stdlib.h>
#include <stdbool.h>

typedef struct {
	long long status;
	char* name;
} A;

typedef struct {
	bool status;
	double value;
	char* name;
	A a;
} Input;

typedef struct {
	bool status;
	double value;
	char* name;
} Output;
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export test
func test(input C.Input) C.Output {
	fmt.Println(input.status)
	fmt.Println(input.value)
	fmt.Println(C.GoString(input.name))
	fmt.Println("=====")
	fmt.Println(input.a.status)
	fmt.Println(C.GoString(input.a.name))

	return C.Output{
		status: C.bool(true),
		value:  C.double(1.0),
		name:   C.CString("test"),
	}
}

//export test_release
func test_release(output C.Output) {
	C.free(unsafe.Pointer(output.name))
}

func main() {}

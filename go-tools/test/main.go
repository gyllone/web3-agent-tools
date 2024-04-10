package main

/*
#cgo CFLAGS: -I../dependencies
#include <test.h>
*/
import "C"
import (
	"fmt"
)

//export test
func test(foo C.Optional_String, bar C.List_Dict_Int, baz C.Param) C.Output {
	if foo.is_some {
		fmt.Println("foo: value = ", C.GoString(foo.value))
	} else {
		fmt.Println("foo: value = None")
	}
	fmt.Println("bar: len = ", int(bar.len))
	fmt.Println("baz: foo:", baz.foo)

	return C.Output{
		status: C.Bool(true),
		err:    C.CString("error"),
		result: C.some_Param(C.Param{
			foo: C.Bool(true),
			bar: C.new_List_Float(0),
			baz: C.new_Dict_Int(0),
		}),
	}
}

//export test_release
func test_release(output C.Output) {
	C.release_Output(output)
}

func main() {}

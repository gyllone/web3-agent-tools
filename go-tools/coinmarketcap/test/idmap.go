package functest

/*
#cgo CFLAGS: -I/home/ecs-user/zjx/web3-agent-tools/go-tools/output
#cgo LDFLAGS: /home/ecs-user/zjx/web3-agent-tools/go-tools/output/idmap.so
#include "idmap.h"
*/
import "C"
import (
	"testing"
	"unsafe"
)

func queryIdMap(t *testing.T) {
	// 转换 Go 字符串为 C 字符串，记得释放内存
	listingStatusValue := C.CString("")
	defer C.free(unsafe.Pointer(listingStatusValue))

	// 创建 C.OptionalStr 结构体
	var listingStatus C.OptionalStr
	listingStatus.is_some = C.bool(false) // 假设这个值总是有的
	listingStatus.value = listingStatusValue

	// 创建 C.OptionalInt 结构体
	var start C.OptionalInt
	start.is_some = C.bool(false) // 假设 start 总是有值

	var limit C.OptionalInt
	limit.is_some = C.bool(true) // 假设 limit 总是有值
	limit.value = C.longlong(5)

	sortValue := C.CString("")
	defer C.free(unsafe.Pointer(sortValue))

	var sort C.OptionalStr
	sort.is_some = C.bool(false)

	var symbol C.OptionalStr
	symbol.is_some = C.bool(false)

	var aux C.OptionalStr
	aux.is_some = C.bool(false)

	// 调用 C 函数
	result := C.query_id_map(&listingStatus, &start, &limit, &sort, &symbol, &aux)

	defer C.query_id_map_release(result)

	// 校验结果和处理
	if bool(result.is_fail) {
		t.Errorf("query_id_map failed: %s", C.GoString(result.error_message))
	}

	// 打印Keys和Values
	// idMaps := (*[1 << 10]*C.IdMap)(unsafe.Pointer(result.idMaps.idMaps))[:result.idMaps.len:result.idMaps.len]
	// for _, idMap := range idMaps {
	// 	keys := (*[1 << 10]*C.char)(unsafe.Pointer(idMap.keys))[:idMap.len:idMap.len]
	// 	values := (*[1 << 10]*C.char)(unsafe.Pointer(idMap.values))[:idMap.len:idMap.len]
	// 	for i := range keys {
	// 		fmt.Printf("Key: %s, Value: %s\n", C.GoString(keys[i]), C.GoString(values[i]))
	// 	}
	// }

	keys := (*[1 << 10]*C.char)(unsafe.Pointer(result.idMaps.keys))[:result.idMaps.len:result.idMaps.len]
	values := (*[1 << 10]*C.char)(unsafe.Pointer(result.idMaps.values))[:result.idMaps.len:result.idMaps.len]
	for i := range keys {
		t.Logf("Key: %s, Value: %s\n", C.GoString(keys[i]), C.GoString(values[i]))
	}
}

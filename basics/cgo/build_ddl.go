package main

/*
#cgo CFLAGS: -I.
#include <stdlib.h>
#include "process.h"
#include "process.c"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	id := 1001
	score := 98.5

	name := C.CString("Alice")
	defer C.free(unsafe.Pointer(name))

	nums := []int{1, 2, 3, 4, 5}
	cNums := C.malloc(C.size_t(len(nums)) * C.size_t(unsafe.Sizeof(C.int(0))))
	defer C.free(cNums)

	cNumsSlice := (*[1 << 30]C.int)(cNums)[:len(nums):len(nums)]
	for i, v := range nums {
		cNumsSlice[i] = C.int(v)
	}

	result := C.process(
		C.int(id),
		C.double(score),
		name,
		(*C.int)(cNums),
		C.int(len(nums)),
	)

	if result == nil {
		panic("process failed")
	}

	goResult := (*C.Result)(result)
	fmt.Println("sum:", int(goResult.sum))
	fmt.Println("avg:", float64(goResult.avg))
	fmt.Println("message:", C.GoString(goResult.message))

	C.free_result(result)
}

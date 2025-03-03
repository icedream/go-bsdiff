// Copyright (c) 2017-2025 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

//go:build cgo
// +build cgo

package native

/*
#include <stdint.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

var (
	writers = writerTable{}
	readers = readerTable{}
)

func bytesToUint8PtrAndSize(bytes []byte) (ptr *C.uint8_t, size C.int64_t) {
	ptr = (*C.uint8_t)(unsafe.Pointer(&bytes[0]))
	size = C.int64_t(int64(len(bytes)))
	return
}

func cPtrToSlice(ptr *byte, size int) []byte {
	var slice []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = size
	sliceHeader.Len = size
	sliceHeader.Data = uintptr(unsafe.Pointer(ptr))

	return slice
}

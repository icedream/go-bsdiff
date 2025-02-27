// Copyright (c) 2017-2025 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package native

/*
#include "bsdiff.h"
*/
import "C"
import "unsafe"

//export cgo_write_buffer
func cgo_write_buffer(
	bufferIndex C.int,
	dataPtr unsafe.Pointer,
	size C.int,
) C.int {
	buffer := writers.Get(int(bufferIndex))
	errCode := 0
	if _, err := buffer.Write(cPtrToSlice(dataPtr, int(size))); err != nil {
		errCode = 1
	}
	return C.int(errCode)
}

// Copyright (c) 2017-2025 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

//go:build cgo
// +build cgo

package native

/*
#cgo CFLAGS: -I../../bsdiff

#include "bspatch.h"
#include "cgo.h"
*/
import "C"

import (
	"errors"
	"io"
	"strconv"
)

func Patch(oldbytes, newbytes []byte, patch io.Reader) (err error) {
	oldptr, oldsize := bytesToUint8PtrAndSize(oldbytes)
	newptr, newsize := bytesToUint8PtrAndSize(newbytes)

	bufferIndex := readers.Add(patch)
	defer readers.Free(bufferIndex)

	errCode := int(C.bspatch_cgo(oldptr, oldsize, newptr, newsize, C.int(bufferIndex)))
	if errCode != 0 {
		err = errors.New("bspatch failed with code " + strconv.Itoa(errCode))
		return
	}

	return
}

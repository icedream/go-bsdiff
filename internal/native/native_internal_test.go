// Copyright (c) 2017-2025 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

//go:build cgo && go1.20
// +build cgo,go1.20

package native

import (
	"crypto/rand"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const testBytesLen = 512 * 1024

func makeTestBytes() []byte {
	b := make([]byte, testBytesLen)
	_, _ = rand.Read(b)
	return b
}

func Test_bytesToUint8PtrAndSize(t *testing.T) {
	b := makeTestBytes()
	ptr, size := bytesToUint8PtrAndSize(b)
	assert.EqualValues(t, &b[0], ptr)
	assert.EqualValues(t, testBytesLen, size)
}

func Test_cPtrToSlice(t *testing.T) {
	b := makeTestBytes()
	bytePtr := unsafe.SliceData(b)
	b2 := cPtrToSlice(bytePtr, testBytesLen)
	assert.Equal(t, b, b2)
}

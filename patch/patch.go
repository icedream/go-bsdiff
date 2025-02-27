// Copyright (c) 2017-2025 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package patch

import (
	"compress/bzip2"
	"io"
	"io/ioutil"

	"github.com/icedream/go-bsdiff/internal"
	"github.com/icedream/go-bsdiff/internal/native"
)

func Patch(oldReader io.Reader, newWriter io.Writer, patchReader io.Reader) (err error) {
	oldBytes, err := ioutil.ReadAll(oldReader)
	if err != nil {
		return
	}

	newLen, err := internal.ReadHeader(patchReader)
	if err != nil {
		return
	}
	newBytes := make([]byte, newLen)

	// Decompression
	bz2Reader := bzip2.NewReader(patchReader)

	err = native.Patch(oldBytes, newBytes, bz2Reader)
	if err != nil {
		return err
	}

	_, err = newWriter.Write(newBytes)
	return
}

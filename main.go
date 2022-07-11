// Copyright (c) 2017-2022 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package bsdiff

import (
	"io"

	"github.com/icedream/go-bsdiff/diff"
	"github.com/icedream/go-bsdiff/patch"
)

func Diff(oldReader, newReader io.Reader, patchWriter io.Writer) (err error) {
	return diff.Diff(oldReader, newReader, patchWriter)
}

func Patch(oldReader io.Reader, newWriter io.Writer, patchReader io.Reader) (err error) {
	return patch.Patch(oldReader, newWriter, patchReader)
}

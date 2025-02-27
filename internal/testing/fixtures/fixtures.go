// Copyright (c) 2017-2025 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package fixtures

import _ "embed"

// The updated test file.
//
//go:embed new.txt
var New []byte

// The original test file.
//
//go:embed old.txt
var Old []byte

// A fully bsdiff-formatted patch between the original and the updated test file.
//
//go:embed old_to_new_patch.bin
var Patch []byte

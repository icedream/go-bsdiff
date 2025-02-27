// Copyright (c) 2017-2025 Carl Kittelberger.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package patch_test

import (
	"bytes"
	"testing"

	"github.com/icedream/go-bsdiff/internal/testing/fixtures"
	"github.com/icedream/go-bsdiff/patch"
	"github.com/stretchr/testify/require"
)

func TestPatch(t *testing.T) {
	t.Run("succeeds with valid patch", func(t *testing.T) {
		oldReader := bytes.NewReader(fixtures.Old)
		patchReader := bytes.NewReader(fixtures.Patch)
		var buf bytes.Buffer

		require.NoError(t, patch.Patch(oldReader, &buf, patchReader))
		require.Equal(t, fixtures.New, buf.Bytes())
	})

	t.Run("fails with invalid patch", func(t *testing.T) {
		oldReader := bytes.NewReader(fixtures.Old)
		patchReader := bytes.NewReader([]byte("invalid patch"))
		var buf bytes.Buffer

		require.Error(t, patch.Patch(oldReader, &buf, patchReader))
		require.Empty(t, buf.Bytes())
	})

	t.Run("fails with invalid original file", func(t *testing.T) {
		// TODO - bsdiff currently seems to return 0 code...?
		t.Skip()

		oldReader := bytes.NewReader([]byte("bad file"))
		patchReader := bytes.NewReader(fixtures.Patch)
		var buf bytes.Buffer

		require.Error(t, patch.Patch(oldReader, &buf, patchReader))
		require.Empty(t, buf.Bytes())
	})
}

func BenchmarkPatch(b *testing.B) {
	// make sure we start in the middle of the loop from exactly 0 time passed
	b.StopTimer()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		oldReader := bytes.NewReader(fixtures.Old)
		patchReader := bytes.NewReader(fixtures.Patch)
		var buf bytes.Buffer

		// measure actual patch
		b.StartTimer()
		err := patch.Patch(oldReader, &buf, patchReader)
		b.StopTimer()

		require.NoError(b, err)
	}
}

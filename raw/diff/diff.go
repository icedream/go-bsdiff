package diff

import (
	"io"
	"io/ioutil"

	"github.com/icedream/go-bsdiff/internal/native"
)

/*
Diff generates a raw patch from old content that will be read in completely from
oldReader and new content that will be read in completely from newReader and
saves that patch to patchWriter.

It may be helpful to save away the new content size along with the actual
patch as it will be needed in order to reuse the patch.
*/
func Diff(oldReader, newReader io.Reader, patchWriter io.Writer) (err error) {
	oldBytes, err := ioutil.ReadAll(oldReader)
	if err != nil {
		return
	}
	newBytes, err := ioutil.ReadAll(newReader)
	if err != nil {
		return
	}

	return native.Diff(oldBytes, newBytes, patchWriter)
}

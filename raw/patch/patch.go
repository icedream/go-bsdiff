package patch

import (
	"io"
	"io/ioutil"

	"github.com/icedream/go-bsdiff/internal/native"
)

/*
Patch reads a raw patch from patchReader and applies it on top of the old
content which will be read from oldReader, saving the resulting new content to
newWriter.

newSize needs to be exactly the size of the new file that should be generated
from the patch.
*/
func Patch(oldReader io.Reader, newWriter io.Writer, patchReader io.Reader, newSize uint64) (err error) {
	oldBytes, err := ioutil.ReadAll(oldReader)
	if err != nil {
		return
	}

	newBytes := make([]byte, newSize)

	err = native.Patch(oldBytes, newBytes, oldReader)

	newWriter.Write(newBytes)
	return
}

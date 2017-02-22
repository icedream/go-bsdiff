package main

import (
	"log"
	"os"

	"github.com/icedream/go-bsdiff/patch"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cli = kingpin.New("go-bspatch", "Applies binary patches generated using the bsdiff algorithm.")

	argOld   = cli.Arg("old", "The old file.").Required().ExistingFile()
	argNew   = cli.Arg("new", "Where the new file will be written to.").Required().File()
	argPatch = cli.Arg("patch", "The patch file.").Required().ExistingFile()
)

func must(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)
}

func main() {
	kingpin.MustParse(cli.Parse(os.Args[1:]))

	newFile := *argNew
	defer newFile.Close()

	oldFile, err := os.Open(*argOld)
	must(err)
	defer oldFile.Close()

	patchFile, err := os.Open(*argPatch)
	must(err)
	defer patchFile.Close()

	must(patch.Patch(oldFile, newFile, patchFile))
}

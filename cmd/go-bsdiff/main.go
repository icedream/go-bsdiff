package main

import (
	"log"
	"os"

	"github.com/icedream/go-bsdiff/diff"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cli = kingpin.New("go-bsdiff", "Generates binary patches.")

	argOld   = cli.Arg("old", "The old file.").Required().ExistingFile()
	argNew   = cli.Arg("new", "The new file.").Required().ExistingFile()
	argPatch = cli.Arg("patch", "Where to output the patch file.").Required().String()
)

func must(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)
}

func main() {
	kingpin.MustParse(cli.Parse(os.Args[1:]))

	patchFile, err := os.Create(*argPatch)
	must(err)
	defer patchFile.Close()

	oldFile, err := os.Open(*argOld)
	must(err)
	defer oldFile.Close()

	newFile, err := os.Open(*argNew)
	must(err)
	defer newFile.Close()

	must(diff.Diff(oldFile, newFile, patchFile))
}

package main

import (
	"log"
	"os"

	"github.com/icedream/go-bsdiff"
)

func must(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)
}

func main() {
	oldFile, err := os.Open("test.txt")
	must(err)

	newFile, err := os.Open("test_new.txt")
	must(err)

	patchFile, err := os.Create("test.patch")
	must(err)

	must(bsdiff.Diff(oldFile, newFile, patchFile))

	log.Println("Done.")
}

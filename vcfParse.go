package main

import (
	"fmt"
	"githu"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
	"github.com/srynobio/govmc"
	"log"
	"os"
)

func main() {

	fh, err := xopen.Ropen(os.Args[1])
	eCheck(err)
	defer fh.Close()

	rdr, err := vcfgo.NewReader(fh, false)
	eCheck(err)
	defer rdr.Close()

	for {
		variant := rdr.Read()
		if variant == nil {
			break
		}

		fmt.Println(variant.Id())
		// Check for alternate alleles.
		altAllele := variant.Alt()

		if len(altAllele) > 1 {
			log.Panicln("multiallelic variants found, please pre-run vt decomposes.")
		}
	}
}

// ------------------------- //

func eCheck(p error) {
	if p != nil {
		panic(p)
	}
}

// ------------------------- //

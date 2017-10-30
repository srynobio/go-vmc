package main

import (
	//"fmt"
	"fmt"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
	"github.com/srynobio/go-vmc/vmc"
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

		// Check for alternate alleles.
		altAllele := variant.Alt()
		if len(altAllele) > 1 {
			log.Panicln("multiallelic variants found, please pre-run vt decomposes.")
		}

		// set variant line to build vmc
		record := vmc.GetVMCRecord(variant)

		fmt.Println(record.Location)

	}
}

// ------------------------- //

func eCheck(p error) {
	if p != nil {
		panic(p)
	}
}

// ------------------------- //

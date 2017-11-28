package main

import (
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

	// Add VMC INFO to the header.
	rdr.AddInfoToHeader("VMCGSID", "1", "String", "VMC Sequence identifier")
	rdr.AddInfoToHeader("VMCGLID", "1", "String", "VMC Location identifier")
	rdr.AddInfoToHeader("VMCGAID", "1", "String", "VMC Allele identifier")

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
		record := vmc.VMCRecord(variant)

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

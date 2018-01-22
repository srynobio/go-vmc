package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
	"github.com/srynobio/go-vmc/vmc"
)

func main() {

	var args struct {
		VCF       string `help: "VCF file to annotate with VMC digest ids"`
		Fasta     string `help: "Reference fasta file used to create VCF file."`
		NameSpace string `arg: "-nameSpace: namespace used to build VMC digest.`
	}
	arg.MustParse(&args)

	fh, err := xopen.Ropen(args.VCF)
	if err != nil {
		panic("VCF file not given or could not be opened.")
	}
	defer fh.Close()

	rdr, err := vcfgo.NewReader(fh, false)
	if err != nil {
		panic("Could not read given VCF file.")
	}
	defer rdr.Close()

	// Add VMC INFO to the header.
	rdr.AddInfoToHeader("VMCGSID", "1", "String", "VMC Sequence identifier")
	rdr.AddInfoToHeader("VMCGLID", "1", "String", "VMC Location identifier")
	rdr.AddInfoToHeader("VMCGAID", "1", "String", "VMC Allele identifier")

	for {
		variant := rdr.Read()
		if variant == nil {
			break
		}

		// Check for alternate alleles.
		altAllele := variant.Alt()
		if len(altAllele) > 1 {
			panic("multiallelic variant found, please pre-run with vt.")
		}

		fmt.Println("~~~~~ new record ~~~~~~~~")
		record := vmc.VMCMarshal(variant, "VMC")

		fmt.Println(vmc.AlleleID(record))
		fmt.Println(vmc.LocationID(record))

	}
}

// ------------------------- //
